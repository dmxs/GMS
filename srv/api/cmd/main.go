package main

import (
	"GMS/pkg/logger"
	"GMS/srv/api/router"
	"log"
)

func init() {
	logger.InitLog()
}

func main() {
	r := router.NewRouter()
	if err := r.Run("127.0.0.1:8080"); err != nil {
		log.Print(err.Error())
	}
}
