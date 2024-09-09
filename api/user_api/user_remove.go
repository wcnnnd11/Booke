package user_api

import (
	"GVB_server/global"
	"GVB_server/models"
	"GVB_server/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (UserApi) UserRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var userList []models.UserModel
	count := global.DB.Find(&userList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("用户不存在", c)
		return
	}

	//事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		//TODO:删除用户消息标，评论标，收藏的文章，发布的文章
		err = global.DB.Model(&userList).Association("Banners").Clear()
		err = global.DB.Delete(&userList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除用户失败", c)
	}

	res.OkWithMessage(fmt.Sprintf("共删除 %d 个用户", count), c)
}
