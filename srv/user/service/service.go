package service

import (
	"GMS/srv/user/conf"
	"GMS/srv/user/dao"
	"GMS/srv/user/proto/role"
	"github.com/micro/go-micro/v2/client"
)

//Service .
type Service struct {
	c   *conf.Config
	dao *dao.Dao
	roleSvc role.RoleService
}

//New
func New(c *conf.Config, client client.Client) (s *Service) {
	s = &Service{
		c:   c,
		dao: dao.New(c),
		roleSvc : role.NewRoleService(c.Remote.Role,client),
	}
	return s
}
