package ctrl

import (
	"ForumPublica/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

//RootIndex index
func RootIndex(c *gin.Context) {
	c.Keys["DeviationsOver"], c.Keys["DeviationsUnder"] = services.DeviationsList()
	c.Keys["Title"] = "Publica"
	c.HTML(http.StatusOK, "root/index.html", c.Keys)
}
