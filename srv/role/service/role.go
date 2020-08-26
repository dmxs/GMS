package service

import (
	"GMS/pkg/common"
	"GMS/pkg/logger"
	"GMS/srv/role/conf"
	"GMS/srv/role/proto/role"
	"context"
	"github.com/micro/go-micro/v2/errors"
)

func (s *Service) Info(ctx context.Context, req *role.InfoReq, reply *role.InfoReply) error {
	info, exist, err := s.dao.RoleGet(req.Id)
	if !exist {
		return errors.New(conf.Conf.Micro.Name,common.MsgRoleNotExist ,common.StatusRoleNotExist)
	}

	if err != nil {
		logger.Error(err.Error())
		return errors.New(conf.Conf.Micro.Name,common.MsgServerError ,common.StatusServerError)
	}

	reply.Id = info.ID
	reply.Name = info.Name

	return nil
}
