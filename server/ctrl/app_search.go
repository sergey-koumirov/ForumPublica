package ctrl

import (
	"ForumPublica/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AppSearchItemType(c *gin.Context) {
	temp := services.SearchItemType(c.Query("term"))
	c.JSON(http.StatusOK, temp)
}

func AppSearchBlueprint(c *gin.Context) {
	temp := services.SearchBlueprint(c.Query("term"))
	c.JSON(http.StatusOK, temp)
}
