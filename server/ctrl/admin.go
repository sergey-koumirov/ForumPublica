package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//AdminIndex index
func AdminIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.html", c.Keys)
}
