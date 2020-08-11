package service

import (
	"GMS/pkg/common"
	"GMS/pkg/logger"
	"GMS/srv/user/model"
	"GMS/srv/user/proto"
	"context"
	"strconv"
)

//Post 添加用户
func (s *Service) Post(ctx context.Context, req *user.PostReq, resp *user.PostResp) (err error) {
	var ok bool
	_, ok, err = s.Dao.UserGetByUsername(req.Username)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	if ok {
		logger.Info("user has registered")
		resp.Error = &user.Error{Code: common.StatusUserHasRegistered, Msg: common.UserHasRegisteredMsg}
		return
	}

	password := Sum256(req.Password)

	user, err := s.Dao.UserPost(req.Username, password)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	logger.Info("create user :" + user.Username)
	return
}

//CheckPassword .
func (s *Service) CheckPassword(ctx context.Context, req *user.CheckPasswordReq, resp *user.CheckPasswordResp) (err error) {
	password := Sum256(req.Password)

	userM, ok, err := s.Dao.UserCheckPassword(req.Username, password)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	resp.Result = ok
	if ok {
		resp.User = &user.UserItem{
			Id:       parseID2Str(userM.ID),
			Username: userM.Username,
			Img:      userM.Image,
		}
	}
	return
}

//GetList 获取用户列表
func (s *Service) GetList(ctx context.Context, req *user.GetListReq, resp *user.GetListResp) (err error) {
	userMList, err := s.Dao.UserGetList(req.Name, req.Ids)
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
	userM, exist, err := s.Dao.UserGet(req.Id)
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
			Username: userM.Username,
			Img:      userM.Image,
		}
	}

	return
}

func dtoUserMList2PbUserItem(userMList []*model.User) (userItems []*user.UserItem) {

	for _, userM := range userMList {
		item := new(user.UserItem)
		item.Id = parseID2Str(userM.ID)
		item.Username = userM.Username
		item.Img = userM.Image

		userItems = append(userItems, item)
	}
	return
}

func parseID2Str(id uint) (str string) {
	str = strconv.FormatUint(uint64(id), 10)
	return
}
