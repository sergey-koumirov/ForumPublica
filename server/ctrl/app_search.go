package ctrl

import (
	"ForumPublica/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AppSearchItemType search
func AppSearchItemType(c *gin.Context) {
	temp := services.SearchItemType(c.Query("term"), c.Param("filter"))
	c.JSON(http.StatusOK, temp)
}
