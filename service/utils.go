package service

import (
	"github/zhblogs/backend/middleware"
	"github/zhblogs/backend/provider"
	"github/zhblogs/backend/utils/status"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func getProvider(c *gin.Context) provider.Provider {
	p, exist := c.Get(middleware.ContextProviderName)
	if !exist {
		err := errors.New("got provider failed")
		logrus.Panicf("%+v", err)
	}

	return p.(provider.Provider)
}

func succWithJSON(c *gin.Context, data interface{}) {
	c.JSON(status.OK, CommonResponse{
		Status: true,
		Data:   data,
	})
}

func failedWithError(c *gin.Context, err error) {
	logrus.Errorf("%+v", errors.Wrap(err, "internal error"))

	c.JSON(status.OK, CommonResponse{
		Status:  false,
		Message: err.Error(),
	})
}
