package router

import (
	"github/zhblogs/backend/service"

	"github.com/gin-gonic/gin"
)

const (
	blogsPrefix        = "blogs"
	featureBlogsPrefix = "featured-blogs"
	randomBlogsPrefix  = "random-blogs"
)

func BindingRouter(g *gin.Engine) {
	{
		randomBlogs := g.Group(blogsPrefix)
		randomBlogs.GET("", service.GetBlogs)
	}
	{
		featureBlogs := g.Group(featureBlogsPrefix)
		featureBlogs.GET("", service.GetFeatureBlogs)
	}
	{
		randomBlogs := g.Group(randomBlogsPrefix)
		randomBlogs.GET("", service.GetRandomBlogs)
	}
}
