package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.html", c.Keys)
}
