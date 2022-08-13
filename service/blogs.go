package service

import (
	"github/zhblogs/backend/provider/types"
	"github/zhblogs/backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GetFeatureBlogs handle GET /feature-blogs
func GetFeatureBlogs(c *gin.Context) {
	p := getProvider(c)
	res, err := p.GetFeatureBlogs()
	if err != nil {
		failedWithError(c, err)
		return
	}

	succWithJSON(c, res)
}

// GetRandomBlogs handle GET /random-blogs
func GetRandomBlogs(c *gin.Context) {
	tagsStr := c.Query("tags")
	countStr := c.DefaultQuery("count", "10")

	tags := strings.Split(tagsStr, ",")
	count, err := utils.ParseInt(countStr, 10)
	if err != nil {
		logrus.Warningf("get query count failed, use default value, %v", err)
	}

	p := getProvider(c)
	res, err := p.GetRandomBlogs(tags, count)
	if err != nil {
		failedWithError(c, err)
		return
	}

	succWithJSON(c, res)
}

type GetBlogsResponse struct {
	Total int                `json:"total"`
	Blogs []types.BlogSimple `json:"blogs"`
}

// GetBlogs handle GET /blogs
func GetBlogs(c *gin.Context) {
	search := c.Query("search")
	tagsStr := c.Query("tags")
	statusStr := c.DefaultQuery("status", "1")
	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "10")

	tags := strings.Split(tagsStr, ",")
	statusInt, _ := utils.ParseInt(statusStr, 1)
	offset, _ := utils.ParseInt(offsetStr, 0)
	limit, _ := utils.ParseInt(limitStr, 10)

	status := types.BlogStatus(statusInt)
	if !status.Vaild() {
		status = types.BlogStatusEnabled
	}

	p := getProvider(c)
	total, blogs, err := p.GetBlogs(search, tags, status, offset, limit)
	if err != nil {
		failedWithError(c, err)
		return
	}

	succWithJSON(c, GetBlogsResponse{
		Total: total,
		Blogs: blogs,
	})
}
