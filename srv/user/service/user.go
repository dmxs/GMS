package service

import (
	"GMS/pkg/common"
	"GMS/pkg/logger"
	"GMS/srv/user/model"
	"GMS/srv/user/proto/role"
	"GMS/srv/user/proto/user"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/micro/go-micro/v2/errors"
)

func (s *Service) GetRole(ctx context.Context,id int64) (re model.RoleReply, err error){
	//rpc
	var reply *role.InfoReply
	if reply, err = s.roleSvc.Info( ctx, &role.InfoReq{ Id:id } ); err != nil {
		return
	}

	re.ID = reply.Id
	re.Name = reply.Name

	return
}

//Post 添加用户
func (s *Service) Post(ctx context.Context, req *user.PostReq, resp *user.PostReply) error {
	if _, ok, err := s.dao.UserGetByUsername(req.Username); err != nil {
		logger.Error(err.Error())
		return err
	} else if ok {
		logger.Info("user has registered")
		return errors.New(s.c.Micro.Name,common.MsgUserHasRegistered,common.StatusUserHasRegistered)
	}

	u, err := s.dao.UserPost(req.Username, Sum256(req.Password))
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	logger.Info("create user :" + u.UserName)

	resp.Id  = u.ID

	return nil
}

// Sum256 sha256加密
func Sum256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
