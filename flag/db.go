package flag

import (
	"GVB_server/global"
	"GVB_server/models"
	"GVB_server/plugins/log_stash"
	"GVB_server/plugins/log_stash_v2"
)

func Makemigrations() {
	var err error
	//global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	// 生成四张表的表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.BannerModel{},
			&models.TagModel{},
			&models.MessageModel{},
			&models.AdvertModel{},
			&models.UserModel{},
			&models.CommentModel{},
			//&models.ArticleModel{},
			&models.UserCollectModel{},
			&models.MenuModel{},
			&models.MenuBannerModel{},
			&models.FadeBackModel{},
			&models.LoginDataModel{},
			&models.ChatModel{},
			&log_stash.LogStashModel{},
			&log_stash_v2.LogModel{},
		)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ success ] 生成数据库表结构成功！")
}
