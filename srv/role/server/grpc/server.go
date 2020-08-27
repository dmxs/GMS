package grpc

import (
	"GMS/srv/role/model"
	"GMS/srv/role/proto/role"
	"GMS/srv/role/service"
	"context"
)

type server struct {
	s *service.Service
}

//Post 添加用户
func (s *server) Info(ctx context.Context, req *role.InfoReq, resp *role.InfoReply) (err error) {
	var info model.Role
	if info ,err = s.s.Info(ctx,req.Id); err != nil {
		return err
	}

	resp.Id = info.ID
	resp.Name = info.Name

	return nil
}
