package http

import (
	"GMS/pkg/common"
	"GMS/pkg/convert"
	"GMS/pkg/logger"
	"GMS/srv/user/dao"
	"GMS/srv/user/model"
	"github.com/gin-gonic/gin"
)

 func getByID(c *gin.Context) {
 	//param
 	id := c.GetInt64("id")

 	//查询
 	u, exist, err := dao.DAO.UserGet(id)
	if err != nil {
		logger.Error("查询异常 : " + err.Error())
		common.ResFailCode(c,common.StatusServerError,common.MsgServerError)
		return
	}

 	if !exist {
 		logger.Info("用户不存在")
 		common.ResFailCode(c,common.StatusUserNotExist,common.MsgUserNotExist)
 		return
	}

	common.ResSuccess(c,u)
 }

func getRole(c *gin.Context) {
	//param
	type request struct {
		ID int64 `json:"id"`
	}

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		common.ResFailCode(c,common.StatusInvalidParam,common.MsgInvalidParam)
		return
	}

	logger.Info(convert.ToString(c.GetInt("id")))

	//查询
	u, exist, err := dao.DAO.UserGet(req.ID)
	if err != nil {
		logger.Error("查询异常 : " + err.Error())
		common.ResFailCode(c,common.StatusServerError,common.MsgServerError)
		return
	}

	if !exist {
		logger.Info("用户不存在")
		common.ResFailCode(c,common.StatusUserNotExist,common.MsgUserNotExist)
		return
	}

	var reply model.RoleReply
	if reply, err = svr.GetRole(c,u.ID); err != nil {
		common.ResFailError(c,err)
		return
	}

	common.ResSuccess(c,reply)
}
