package middleware

import (
	mock "github/zhblogs/backend/provider/mock_provider"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	ContextProviderName = "Provider"
)

func ProviderHandler() gin.HandlerFunc {
	provider := mock.New()

	logrus.Infof("init provider")
	return func(c *gin.Context) {
		logrus.Infof("set provider")
		c.Set(ContextProviderName, provider)

		c.Next()
	}
}
