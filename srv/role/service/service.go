package service

import (
	"GMS/srv/role/conf"
	"GMS/srv/role/dao"
)

//Service .
type Service struct {
	c   *conf.Config
	dao *dao.Dao
}

//New
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:   c,
		dao: dao.New(c),
	}
	return s
}
