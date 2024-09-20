package common

import (
	"GVB_server/global"
	"GVB_server/models"
	"fmt"
	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
	Likes []string // 模糊匹配的字段
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {

	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})

	}
	if option.Sort == "" {
		option.Sort = "created_at desc" //默认按照时间往前排
	}
	DB = DB.Where(model)

	for index, column := range option.Likes {
		if index == 0 {
			DB.Where(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
			continue
		}
		DB.Or(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
	}

	count = DB.Find(&list).RowsAffected
	//这里的query受上面影响，需要手动复位
	query := DB.Where(model)
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}

	/*
		Limit有问题，应该是option.Limit,但是option.Limit=0,解决方法推荐2种
		 判断limit=0的情况，并重新赋值为1
		 或者在每次查询的时候手动设置limit的值
		 不推荐的方法：更换Gorm的版本至 1.24.5
	*/
	// 判断limit = 0 的解决方法
	pageLimit := option.Limit
	if pageLimit == 0 {
		pageLimit = -1
	}
	err = query.Limit(pageLimit).Offset(offset).Order(option.Sort).Find(&list).Error

	return list, count, err
}
