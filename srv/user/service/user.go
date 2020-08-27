package service

import (
	"GMS/pkg/common"
	"GMS/pkg/logger"
	"GMS/srv/user/model"
	"GMS/srv/user/proto/role"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/micro/go-micro/v2/errors"
)

func (s *Service) GetRole(ctx context.Context, roleID int64) (re model.RoleReply, err error){
	//rpc
	var reply *role.InfoReply
	if reply, err = s.roleSvc.Info( ctx, &role.InfoReq{ Id: roleID } ); err != nil {
		return
	}

	re.ID = reply.Id
	re.Name = reply.Name

	return
}

func (s *Service) GetByID(ctx context.Context, id int64) (info *model.User, err error) {
	var exist bool
	if info, exist, err = s.dao.UserGetByID(id); err != nil {
		logger.Error(err.Error())
		return
	} else if !exist {
		err = errors.New(s.c.Micro.Name, common.MsgUserNotExist, common.StatusUserNotExist)
	}

	return
}

//Post 添加用户
func (s *Service) Post(ctx context.Context, username string, password string) (id int64, err error) {
	var exist bool
	if _, exist, err = s.dao.UserGetByUsername(username); err != nil {
		logger.Error(err.Error())
		return
	} else if exist {
		err = errors.New(s.c.Micro.Name, common.MsgUserHasRegistered, common.StatusUserHasRegistered)
	}

	var info *model.User
	if info, err = s.dao.UserPost(username, Sum256(password)); err != nil {
		err = errors.New(s.c.Micro.Name, common.MsgServerError, common.StatusServerError)
		return
	}

	logger.Info("create user :" + info.UserName)

	return info.ID,nil
}

// Sum256 sha256加密
func Sum256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
