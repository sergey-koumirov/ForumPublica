package ctrl

import (
	"ForumPublica/server/middleware"
	"ForumPublica/server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func page(c *gin.Context) int64 {
	page, errParse := strconv.ParseInt(c.Query("page"), 10, 64)
	if errParse != nil {
		page = 1
	}
	return page
}

func user(c *gin.Context) models.User {
	raw, _ := c.Get(middleware.USER)

	if raw == nil {
		return models.User{}
	}

	return raw.(models.User)
}
