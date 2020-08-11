package model

import (
	"time"
)

type Base struct {
	ID       uint       `gorm:"column:id; primary_key; auto_increment;"`
	CreateAt *time.Time `gorm:"column:created_at; not null; default:CURRENT_TIMESTAMP"`
	UpdateAt *time.Time `gorm:"column:updated_at; not null; default:CURRENT_TIMESTAMP"`
}

//ModelUser .
type User struct {
	Base
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Image    string `gorm:"column:image"`
}

//TableName .
func (a *User) TableName() string {
	return "user"
}
