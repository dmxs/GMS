package dao

import (
	"GMS/pkg/gormplus"
	"GMS/pkg/logger"
	"GMS/srv/role/conf"
)

var (
	DAO *Dao
)

type Dao struct {
	DB *gormplus.DB
}

func (d *Dao) Ping() (err error) {
	err = d.DB.DB.DB().Ping()
	return
}

//Init dao
func New(c *conf.Config) *Dao {
	DAO = &Dao{
		DB: gormplus.New(&gormplus.Config{
			Debug: c.MySQL.Debug,
			DSN:   c.MySQL.DSN(),
		}),
	}

	if err := DAO.Ping(); err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info("数据库初始化成功")
	}

	return DAO
}
