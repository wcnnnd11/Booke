package models

import "time"

// UserCollectModel 自定义第三张表  记录用户什么时候收藏了什么文章
type UserCollectModel struct {
	UserID    uint      `gorm:"primaryKey;autoIncrement:false"` // 联合主键之一
	UserModel UserModel `gorm:"foreignKey:UserID"`
	ArticleID string    `gorm:"primaryKey;size:32"` // 联合主键之一
	CreatedAt time.Time
}
