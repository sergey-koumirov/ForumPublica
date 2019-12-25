package ctrl

import (
	"ForumPublica/server/services"
	"net/http"
	"strconv"

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
	cid, _ := strconv.ParseInt(c.Query("cid"), 10, 64)

	temp, err := services.SearchLocation(u.ID, cid, c.Query("term"), c.Query("filter"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, temp)
	}

}
