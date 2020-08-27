package service

import (
	"GMS/pkg/common"
	"GMS/pkg/logger"
	"GMS/srv/role/conf"
	"GMS/srv/role/model"
	"context"
	"github.com/micro/go-micro/v2/errors"
)

func (s *Service) Info(ctx context.Context, id int64) (info model.Role, err error) {
	var exist bool

	if info, exist, err = s.dao.RoleGet(id); err != nil {
		logger.Error(err.Error())
		err = errors.New(conf.Conf.Micro.Name,common.MsgServerError ,common.StatusServerError)
		return
	} else if !exist {
		err = errors.New(conf.Conf.Micro.Name,common.MsgRoleNotExist ,common.StatusRoleNotExist)
		return
	}

	return
}
