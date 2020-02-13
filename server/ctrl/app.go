package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//AppIndex root page
func AppIndex(c *gin.Context) {
	c.Keys["Title"] = "Index"
	c.HTML(http.StatusOK, "app/index.html", c.Keys)
}
