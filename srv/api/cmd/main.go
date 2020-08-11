package cmd

import (
	"GMS/srv/api/router"
	"log"
)

func main() {
	r := router.NewRouter()
	if err := r.Run("127.0.0.1:8080"); err != nil {
		log.Print(err.Error())
	}
}
