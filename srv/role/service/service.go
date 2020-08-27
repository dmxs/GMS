package service

import (
	"GMS/srv/role/conf"
	"GMS/srv/role/dao"
	"github.com/micro/go-micro/v2/client"
)

//Service .
type Service struct {
	c   *conf.Config
	dao *dao.Dao
}

//New
func New(c *conf.Config, client client.Client) (s *Service) {
	s = &Service{
		c:   c,
		dao: dao.New(c),
	}
	return s
}
