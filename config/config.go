package config

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	Addr string
	Port int
	Prod bool
}

// Global config
var Global = &Config{}

func init() {
	viper.SetDefault("ADDR", "127.0.0.1")
	viper.SetDefault("PORT", 8000)
	viper.SetDefault("PROD", false)

	viper.AutomaticEnv()

	err := viper.Unmarshal(Global)
	if err != nil {
		logrus.Panicf("init environment failed, due to %s", err)
	}

	if !Global.Prod {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}
