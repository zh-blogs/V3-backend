package main

import (
	"fmt"
	"github/zhblogs/backend/config"
	"github/zhblogs/backend/middleware"
	"github/zhblogs/backend/router"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	Gin := gin.Default()
	Gin.Use(middleware.List...)
	router.BindingRouter(Gin)

	addr := fmt.Sprintf("%s:%d", config.Global.Addr, config.Global.Port)
	logrus.Printf("Server start at %s", addr)

	err := Gin.Run(addr)
	if err != nil {
		logrus.Infof("Server stoped due to %s", err)
	}
}
