package service

import (
	"GMS/pkg/common"
	"GMS/pkg/logger"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	microhttp "github.com/micro/go-plugins/client/http/v2"
	"io/ioutil"
	"net/http"
	"path"
	"regexp"
	"strings"
)

var (
	versionRe = regexp.MustCompilePOSIX("^v[0-9]+$")
)

func RPC(c *gin.Context) {
	//解析
	service, method := pathToReceiver(common.PrefixNamespace,common.SuffixNamespace, c.Request.URL.Path)
	if service == "" || method == "" {
		logger.Error("service or method not found")
		common.ResFailCode(c, common.StatusServerError, common.MsgServerError)
		return
	}

	//请求参数
	var br []byte
	param := make(map[string]string)
	if c.Request.Method == "GET" {
		//get
		logger.Info("get请求")
		_ = c.Request.ParseForm()
		for k, v := range c.Request.Form {
			param[k] = v[0]
		}
		br, _ = json.Marshal(param)
	} else {
		//other
		logger.Info("非get请求")
		br, _ = ioutil.ReadAll(c.Request.Body)
	}

	//request
	request := json.RawMessage(br)

	//response
	var response json.RawMessage

	//log
	logger.InfoKV("RPC TRANSFER", logger.KV{"service": service, "method": method, "request": request})

	//NewRequest
	microClient := microhttp.NewClient(client.ContentType("application/json"))

	req := microClient.NewRequest(service, method, request)

	if err := microClient.Call(context.Background(), req, &response); err != nil {
		logger.Error("api service call failed:" + err.Error())
		common.ResFailCode(c, common.StatusServerError, common.MsgServerError)
		return
	}

	c.JSON(http.StatusOK,response)
}

func pathToReceiver(prefix string, suffix string, p string) (string, string) {
	p = strings.TrimPrefix(path.Clean(p), "/")
	parts := strings.Split(p, "/")

	// If we've got two or less parts
	// Use first part as service
	// Use all parts as method
	if len(parts) <= 2 {
		service := prefix + strings.Join(parts[:len(parts)-1], ".") + suffix
		method := "/" + strings.Join(parts, "/")
		return service, method
	}

	// Treat /v[0-9]+ as version where we have 3 parts
	// /v1/foo/bar => service: v1.foo method: Foo.bar
	if len(parts) == 3 && versionRe.Match([]byte(parts[0])) {
		service := prefix + strings.Join(parts[:len(parts)-1], ".") + suffix
		method := "/" + strings.Join(parts[len(parts)-2:], "/")
		return service, method
	}

	// Service is everything minus last two parts
	// Method is the last two parts
	service := prefix + strings.Join(parts[:len(parts)-2], ".") + suffix
	method := "/" + strings.Join(parts[len(parts)-2:], "/")
	return service, method
}
