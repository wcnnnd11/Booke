package common

import (
	"GVB_server/global"
	"GVB_server/models"
	"fmt"
	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo          // 分页查询
	Likes           []string // 需要模糊匹配的字段列表
	Debug           bool     // 是否打印sql
	Where           *gorm.DB // 额外的查询
	Preload         []string // 预加载的字段列表
	Role            int      // 新增的角色字段
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	query := global.DB.Where(model)
	if option.Debug {
		query = query.Debug()
	}

	// 处理 `role` 条件
	if option.Role != 0 {
		query = query.Where("role = ?", option.Role)
	}

	if option.Sort == "" {
		option.Sort = "created_at desc"
	}
	if option.Limit == 0 {
		option.Limit = 10
	}
	if option.Where != nil {
		query.Where(option.Where)
	}

	if option.Key != "" {
		likeQuery := global.DB.Where("")
		for index, column := range option.Likes {
			if index == 0 {
				likeQuery.Where(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
			} else {
				likeQuery.Or(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
			}
		}
		query = query.Where(likeQuery)
	}

	count = query.Find(&list).RowsAffected

	for _, preload := range option.Preload {
		query = query.Preload(preload)
	}

	offset := (option.Page - 1) * option.Limit

	err = query.Limit(option.Limit).
		Offset(offset).
		Order(option.Sort).Find(&list).Error

	return
}
