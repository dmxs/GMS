package dao

import (
	"GMS/srv/user/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

// UserPost create user
func (d *Dao) UserPost(username string, password string) (user *model.User, err error) {
	user = new(model.User)
	user.UserName = username
	user.Password = password
	err = d.DB.Create(user).Error
	return
}

//UserGetByUsername get user by name
func (d *Dao) UserGetByUsername(username string) (user *model.User, ok bool, err error) {
	user = new(model.User)
	db := d.DB.Where("username = ?", username)
	ok, err = d.DB.FindOne(db, user)
	return
}

//UserCheckPassword .
func (d *Dao) UserCheckPassword(username, password string) (user *model.User, ok bool, err error) {
	user = new(model.User)
	db := d.DB.Where("username = ?", username).Where("password = ?", password)
	ok, err = d.DB.FindOne(db, user)
	return
}

//UserGetList .
func (d *Dao) UserGetList(name string, ids []string) (user []*model.User, err error) {
	var db *gorm.DB

	db = d.DB.Where("username LIKE ?", fmt.Sprintf("%s%s%s", "%", name, "%"))

	if len(ids) != 0 {
		db = d.DB.Where("id IN (?)", ids)
	}

	err = db.Find(&user).Error

	return
}

//UserGetByID .
func (d *Dao) UserGetByID(id int64) (user *model.User, exist bool, err error) {
	user = new(model.User)
	db := d.DB.Where("id = ?", id)
	exist, err = d.DB.FindOne(db, user)

	return
}
