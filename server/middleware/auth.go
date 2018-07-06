package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {

	value, _ := c.Get("user")

	if value == nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
		c.Abort()
		return
	}

}
