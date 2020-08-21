package model

import (
	"time"
)

type Base struct {
	ID       uint       `json:"id"         gorm:"column:id; primary_key; auto_increment;"`
	CreateAt *time.Time `json:"createTime" gorm:"column:created_at; not null; default:CURRENT_TIMESTAMP"`
	UpdateAt *time.Time `json:"updateTime" gorm:"column:updated_at; not null; default:CURRENT_TIMESTAMP"`
}

//ModelUser .
type User struct {
	Base
	UserName string `json:"userName" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	Image    string `json:"image"    gorm:"column:image"`
}

//TableName .
func (a *User) TableName() string {
	return "user"
}
