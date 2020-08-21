package service

import (
	"GMS/pkg/common"
	"GMS/pkg/convert"
	"GMS/pkg/logger"
	"GMS/srv/user/model"
	"GMS/srv/user/proto"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

//Post 添加用户
func (s *Service) Post(ctx context.Context, req *user.PostReq, resp *user.PostResp) error {
	if _, ok, err := s.dao.UserGetByUsername(req.Username); err != nil {
		logger.Error(err.Error())
		return err
	} else if ok {
		logger.Info("user has registered")
		resp.Error = &user.Error{Code: common.StatusUserHasRegistered, Msg: common.UserHasRegisteredMsg}
		return nil
	}

	password := Sum256(req.Password)

	u, err := s.dao.UserPost(req.Username, password)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	logger.Info("create user :" + u.UserName)
	return nil
}

//CheckPassword .
func (s *Service) CheckPassword(ctx context.Context, req *user.CheckPasswordReq, resp *user.CheckPasswordResp) error {
	password := Sum256(req.Password)

	u, ok, err := s.dao.UserCheckPassword(req.Username, password)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	if !ok {
		resp.Result = false
		resp.Error = &user.Error{Code: common.StatusServerError, Msg: common.MsgServerError}
		return nil
	}

	resp.Result = true
	resp.User = &user.UserItem {
		Id:       parseID2Str(u.ID),
		Username: u.UserName,
		Img:      u.Image,
	}
	return nil
}

//GetList 获取用户列表
func (s *Service) GetList(ctx context.Context, req *user.GetListReq, resp *user.GetListResp) (err error) {
	userMList, err := s.dao.UserGetList(req.Name, req.Ids)
	if err != nil {
		logger.Error(err.Error())
		resp.Error = &user.Error{
			Code: 500,
			Msg:  err.Error(),
		}
		return
	}

	resp.List = dtoUserMList2PbUserItem(userMList)
	return
}

//Get 获取当前用户
func (s *Service) Get(ctx context.Context, req *user.GetReq, resp *user.GetResp) (err error) {
	userM, exist, err := s.dao.UserGet(convert.ToInt(req.Id))
	if err != nil {
		logger.Error(err.Error())
		resp.Error = &user.Error{
			Code: 500,
			Msg:  err.Error(),
		}
		return
	}

	resp.Result = exist
	if resp.Result {
		resp.User = &user.UserItem{
			Id:       parseID2Str(userM.ID),
			Username: userM.UserName,
			Img:      userM.Image,
		}
	}

	return
}

func dtoUserMList2PbUserItem(userMList []*model.User) (userItems []*user.UserItem) {
	for _, userM := range userMList {
		item := new(user.UserItem)
		item.Id = parseID2Str(userM.ID)
		item.Username = userM.UserName
		item.Img = userM.Image

		userItems = append(userItems, item)
	}
	return
}

func parseID2Str(id uint) (str string) {
	str = strconv.FormatUint(uint64(id), 10)
	return
}

// Sum256 sha256加密
func Sum256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
