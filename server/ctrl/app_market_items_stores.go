package ctrl

import (
	"ForumPublica/server/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//AppMarketItemsStoresAdd add market item store
func AppMarketItemsStoresAdd(c *gin.Context) {
	u := user(c)

	params := services.LocationParams{}
	c.BindJSON(&params)

	miID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	services.MarketItemsStoresCreate(u.ID, miID, params)

	p := page(c)

	list := services.MarketItemsList(u.ID, p)

	c.JSON(http.StatusOK, list)
}

//AppMarketItemsStoresDelete delete market item store
func AppMarketItemsStoresDelete(c *gin.Context) {
	u := user(c)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	sid, _ := strconv.ParseInt(c.Param("sid"), 10, 64)

	services.MarketItemsStoresDelete(u.ID, id, sid)

	p := page(c)
	list := services.MarketItemsList(u.ID, p)
	c.JSON(http.StatusOK, list)
}
