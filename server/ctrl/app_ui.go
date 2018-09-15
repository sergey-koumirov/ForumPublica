package ctrl

import (
	"ForumPublica/server/middleware"
	"ForumPublica/server/models"
	"ForumPublica/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AppUIOpenMarket(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	params := make(map[string]int64)
	c.BindJSON(&params)

	services.UIOpenMarket(user.Id, params)

	c.JSON(http.StatusOK, gin.H{})
}
