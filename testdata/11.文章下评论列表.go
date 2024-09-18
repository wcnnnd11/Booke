package main

import (
	"GVB_server/core"
	"GVB_server/global"
	"GVB_server/models"
)

func main() {
	//读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	//连接数据库
	global.DB = core.InitGorm()
	FindArticleCommentList("onLL_pEBGZzdoGZAtj0R")

}

func FindArticleCommentList(articleID string) {
	// 先把文章下的根评论查出来
	var RootCommentList []*models.CommentModel
	global.DB.Find(&RootCommentList, "article_id = ? and parent_comment_id is null", articleID)
	// 遍历根评论，查根评论下的所有子评论
	for _, model := range RootCommentList {
		var subCommentList []models.CommentModel
		FindSubComment(*model, &subCommentList)
		model.SubComments = subCommentList
	}
}

// FindSubComment 递归查评论下的子评论
func FindSubComment(model models.CommentModel, subCommentList *[]models.CommentModel) {
	global.DB.Preload("SubComments").Take(&model)
	for _, sub := range model.SubComments {
		*subCommentList = append(*subCommentList, sub)
		FindSubComment(sub, subCommentList)
	}
	return

}
