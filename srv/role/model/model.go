package model

import "time"

type Base struct {
	ID       int64      `json:"id"         gorm:"column:id; primary_key; auto_increment;"`
	CreateAt *time.Time `json:"createTime" gorm:"column:created_at; not null; default:CURRENT_TIMESTAMP"`
	UpdateAt *time.Time `json:"updateTime" gorm:"column:updated_at; not null; default:CURRENT_TIMESTAMP"`
}

//ModelUser .
type Role struct {
	Base
	Name string `json:"name" gorm:"column:name"`
}

//TableName .
func (a *Role) TableName() string {
	return "role"
}
