package router

import (
	"GMS/pkg/common"
	"GMS/pkg/logger"
	"GMS/srv/user/dao"
	"github.com/gin-gonic/gin"
)

 func getByID(c *gin.Context) {
 	//param
 	id := c.GetInt("id")

 	//查询
 	u, exist, err := dao.DAO.UserGet(id)
	if err != nil {
		logger.Error("查询异常 : " + err.Error())
		common.ResFailCode(c,common.StatusUserNotExist,common.MsgUserNotExist)
		return
	}

 	if !exist {
 		logger.Info("用户不存在")
 		common.ResFailCode(c,common.StatusUserNotExist,common.MsgUserNotExist)
 		return
	}

	common.ResSuccess(c,u)

 }
