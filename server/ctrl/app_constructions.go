package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AppConstructions(c *gin.Context) {
	c.HTML(http.StatusOK, "app/constructions.html", c.Keys)
}
