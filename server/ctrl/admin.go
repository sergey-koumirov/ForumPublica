package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//AdminIndex index
func AdminIndex(c *gin.Context) {
	c.Keys["Title"] = "Admin: Index"
	c.HTML(http.StatusOK, "admin/index.html", c.Keys)
}
