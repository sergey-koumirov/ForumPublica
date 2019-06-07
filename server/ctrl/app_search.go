package ctrl

import (
	"ForumPublica/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AppSearchItemType search
func AppSearchItemType(c *gin.Context) {
	temp := services.SearchItemType(c.Query("term"), "")
	c.JSON(http.StatusOK, temp)
}

//AppSearchBlueprint search
func AppSearchBlueprint(c *gin.Context) {
	temp := services.SearchItemType(c.Query("term"), "blueprint")
	c.JSON(http.StatusOK, temp)
}

//AppSearchLocation search
func AppSearchLocation(c *gin.Context) {
	u := user(c)
	temp := services.SearchLocation(u.ID, 1099415243, c.Query("term"))
	c.JSON(http.StatusOK, temp)
}
