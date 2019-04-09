package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//RootIndex index
func RootIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "root/index.html", c.Keys)
}
