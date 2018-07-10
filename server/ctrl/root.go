package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "root/index.html", nil)
}
