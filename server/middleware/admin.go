package middleware

import (
	"ForumPublica/server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Admin(c *gin.Context) {

	value, _ := c.Get(USER)

	if value == nil || value.(models.User).Role != "A" {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		c.Abort()
		return
	}

}
