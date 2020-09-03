package http

import (
	"GMS/pkg/common"
	"GMS/pkg/logger"
	"GMS/srv/user/model"
	"github.com/gin-gonic/gin"
)

 func getByID(c *gin.Context) {
	 //param
	 type request struct {
		 ID int64 `json:"id"`
	 }

	 //反序列化
	 var req request
	 if err := c.ShouldBindJSON(&req); err != nil {
		 common.ResFailCode(c,common.StatusInvalidParam,common.MsgInvalidParam)
		 return
	 }

	 //查询
 	info ,err := svr.GetByID(c, req.ID)
 	if err != nil {
		logger.Error(err.Error())
		common.ResFailError(c,err)
		return
	}

	common.ResSuccess(c, info)
 }

func getRole(c *gin.Context) {
	//param
	type request struct {
		ID int64 `json:"id"`
	}

	//反序列化
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		common.ResFailCode(c, common.StatusInvalidParam, common.MsgInvalidParam)
		return
	}

	//查询用户
	info ,err := svr.GetByID(c, req.ID)
	if err != nil {
		logger.Error(err.Error())
		common.ResFailError(c, err)
		return
	}

	//查询角色
	var reply model.RoleReply
	if reply, err = svr.GetRole(c, info.RoleID); err != nil {
		common.ResFailError(c,err)
		return
	}

	common.ResSuccess(c,reply)
}
