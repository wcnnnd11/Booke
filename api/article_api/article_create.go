package article_api

import (
	"GVB_server/global"
	"GVB_server/models"
	"GVB_server/models/res"
	"GVB_server/service/es_ser"
	"GVB_server/utils/jwts"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"math/rand"
	"strings"
	"time"
)

type ArticleRequest struct {
	Title    string   `json:"title" binding:"required" msg:"文章标题必填"`   // 文章标题
	Abstract string   `json:"abstract"`                                // 文章简介
	Content  string   `json:"content" binding:"required" msg:"文章内容必填"` // 文章内容
	Category string   `json:"category"`                                // 文章分类
	Source   string   `json:"source"`                                  // 文章来源
	Link     string   `json:"link"`                                    // 原文链接
	BannerID uint     `json:"banner_id"`                               // 文章封面id
	Tags     []string `json:"tags"`                                    // 文章标签
}

func (ArticleApi) ArticleCreateView(c *gin.Context) {
	var cr ArticleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID
	userNickName := claims.NickName
	// 校验content  xss

	// 处理content
	unsafe := blackfriday.MarkdownCommon([]byte(cr.Content))
	// 是不是有script标签
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	//fmt.Println(doc.Text())
	nodes := doc.Find("script").Nodes
	if len(nodes) > 0 {
		// 有script标签
		doc.Find("script").Remove()
		converter := md.NewConverter("", true, nil)
		html, _ := doc.Html()
		markdown, _ := converter.ConvertString(html)
		cr.Content = markdown
	}
	if cr.Abstract == "" {
		// 汉字的截取不一样
		abs := []rune(doc.Text())
		// 将content转为html，并且过滤xss，以及获取中文内容
		if len(abs) > 100 {
			cr.Abstract = string(abs[:100])
		} else {
			cr.Abstract = string(abs)
		}
	}

	// 不传banner_id,后台就随机去选择一张
	if cr.BannerID == 0 {
		var bannerIDList []uint
		global.DB.Model(models.BannerModel{}).Select("id").Scan(&bannerIDList)
		if len(bannerIDList) == 0 {
			res.FailWithMessage("没有banner数据", c)
			return
		}
		rand.Seed(time.Now().UnixNano())
		cr.BannerID = bannerIDList[rand.Intn(len(bannerIDList))]
	}

	// 查banner_id下的banner_url
	var bannerUrl string
	err = global.DB.Model(models.BannerModel{}).Where("id = ?", cr.BannerID).Select("path").Scan(&bannerUrl).Error
	if err != nil {
		res.FailWithMessage("banner不存在", c)
		return
	}

	// 查用户头像
	var avatar string
	err = global.DB.Model(models.UserModel{}).Where("id = ?", userID).Select("avatar").Scan(&avatar).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05")

	// 判断文章的标题是否存在

	article := models.ArticleModel{
		CreatedAt:    now,
		UpdatedAt:    now,
		Title:        cr.Title,
		Keyword:      cr.Title,
		Abstract:     cr.Abstract,
		Content:      cr.Content,
		UserID:       userID,
		UserNickName: userNickName,
		UserAvatar:   avatar,
		Category:     cr.Category,
		Source:       cr.Source,
		Link:         cr.Link,
		BannerID:     cr.BannerID,
		BannerUrl:    bannerUrl,
		Tags:         cr.Tags,
	}

	if article.ISExistData() {
		res.FailWithMessage("文章已存在", c)
		return
	}

	err = article.Create()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	go es_ser.AsyncArticleByFullText(article.ID, article.Title, article.Content)
	res.OkWithMessage("文章发布成功", c)

}
