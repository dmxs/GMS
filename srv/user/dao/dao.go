package dao

import (
	"GMS/pkg/gormplus"
	"GMS/pkg/logger"
	"GMS/srv/user/conf"
	"GMS/srv/user/model"
	"sync"
)

var (
	dao  *Dao
	lock sync.Mutex
)

//Interface dao
type Interface interface {
	//添加用户
	UserPost(username, password string) (user *model.User, err error)
	//通过name获取用户信息
	UserGetByUsername(username string) (user *model.User, ok bool, err error)
	//校验密码
	UserCheckPassword(username, password string) (user *model.User, ok bool, err error)
	//获取用户列表
	UserGetList(name string, ids []string) (user []*model.User, err error)
	//通过ID获取用户信息
	UserGet(id string) (user *model.User, exist bool, err error)
}

type Dao struct {
	DB *gormplus.DB
}

func (d *Dao) Ping() (err error) {
	err = d.DB.DB.DB().Ping()
	return
}

//Init dao
func Init() {
	lock.Lock()
	defer lock.Unlock()

	if dao != nil {
		return
	}

	dao = &Dao{ DB: newDB()}

	if err := dao.Ping(); err != nil {
		logger.Info(err.Error())
	}

	return
}

//NewDB 返回gorm链接实例
func newDB() *gormplus.DB {
	c := conf.GetGlobalConfig()
	return gormplus.New(&gormplus.Config{
		Debug: c.MySQL.Debug,
		DSN:   c.MySQL.DSN(),
	})
}

//GetDao .
func GetDao() *Dao {
	return dao
}
