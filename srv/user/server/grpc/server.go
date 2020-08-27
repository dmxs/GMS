package grpc

import (
	"GMS/srv/user/proto/user"
	"GMS/srv/user/service"
	"context"
)

type server struct {
	s *service.Service
}

//Post 添加用户
func (s *server) Post(ctx context.Context, req *user.PostReq, resp *user.PostReply) (err error) {
	var id int64
	if 	id ,err = s.s.Post(ctx,req.Username,req.Password); err != nil {
		return err
	}
	resp.Id  = id
	return nil
}
