package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootIndex(c *gin.Context) {
	user, _ := c.Get("user")

	c.HTML(http.StatusOK, "root/index.html", gin.H{"User": user})
}
