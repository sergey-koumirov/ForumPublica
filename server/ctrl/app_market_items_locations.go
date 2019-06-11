package ctrl

import (
	"ForumPublica/server/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//AppMarketItemsLocationsAdd add market item location
func AppMarketItemsLocationsAdd(c *gin.Context) {
	u := user(c)

	params := services.LocationParams{}
	c.BindJSON(&params)

	miID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	services.MarketItemsLocationsCreate(u.ID, miID, params)

	p := page(c)

	list := services.MarketItemsList(u.ID, p)

	c.JSON(http.StatusOK, list)
}

//AppMarketItemsLocationsDelete delete market item
func AppMarketItemsLocationsDelete(c *gin.Context) {
	u := user(c)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	lid, _ := strconv.ParseInt(c.Param("lid"), 10, 64)

	services.MarketItemsLocationsDelete(u.ID, id, lid)

	p := page(c)
	list := services.MarketItemsList(u.ID, p)
	c.JSON(http.StatusOK, list)
}
