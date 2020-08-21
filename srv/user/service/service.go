package service

import (
	"GMS/srv/user/conf"
	"GMS/srv/user/dao"
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
