package http

import (
	"GMS/pkg/convert"
	"GMS/pkg/logger"
	"GMS/srv/api/conf"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"
)

func Init(c *conf.Config) {
	//web server
	service := web.NewService(
		web.Name(conf.Conf.Micro.Name),
		web.Address(":" + convert.ToString(conf.Conf.Http.Port)),
		web.Handler(NewRouter()),
	)

	if err := service.Run(); err != nil {
		logger.Error("start failed : " + err.Error())
	}
}

func NewRouter() *gin.Engine {
	router := gin.Default()

	//router.Use(middleware.CrossOriginMiddleware())
	//router.Use(middleware.UserAuthMiddleware())

	/*
	     Method         = "OPTIONS"                ; Section 9.2
	                    | "GET"                    ; Section 9.3
	                    | "HEAD"                   ; Section 9.4
	                    | "POST"                   ; Section 9.5
	                    | "PUT"                    ; Section 9.6
	                    | "DELETE"                 ; Section 9.7
	                    | "TRACE"                  ; Section 9.8
	                    | "CONNECT"                ; Section 9.9
	                    | extension-method
	   extension-method = token
	     token          = 1*<any CHAR except CTLs or separators>
	*/

	router.Use(API)

	return router
}
