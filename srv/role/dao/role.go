package dao

import "GMS/srv/role/model"

func (d *Dao) RoleGet(id int64) (user model.Role, exist bool, err error) {
	exist, err = d.DB.FindOne(d.DB.Where("id = ?", id), &user)
	return
}
