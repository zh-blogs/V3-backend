package middleware

import (
	"github.com/gin-gonic/gin"
)

var List = []gin.HandlerFunc{
	ProviderHandler(),
}
