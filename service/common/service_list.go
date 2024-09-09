package common

import (
	"GVB_server/global"
	"GVB_server/models"
	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {

	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})

	}
	if option.Sort == "" {
		option.Sort = "created_at desc" //默认按照时间往前排
	}
	query := DB.Where(model)

	count = query.Select("id").Find(&list).RowsAffected
	//这里的query受上面影响，需要手动复位
	query = DB.Where(model)
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}

	//Limit有问题，应该是option.Limit,但是option.Limit=0,所以暂时用1代替
	err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

	return list, count, err
}
