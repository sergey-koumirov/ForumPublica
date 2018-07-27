package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AppIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "app/index.html", c.Keys)
}
