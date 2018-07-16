package ctrl

import (
	"ForumPublica/server/db"
	"ForumPublica/server/middleware"
	"ForumPublica/server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AppIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "app/index.html", c.Keys)
}

func AppChars(c *gin.Context) {

	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	var chars []models.Character
	db.DB.Where("user_id = ?", user.Id).Find(&chars)

	c.Keys["chars"] = chars

	c.HTML(http.StatusOK, "app/chars.html", c.Keys)
}
