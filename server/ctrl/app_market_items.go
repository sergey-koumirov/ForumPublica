package ctrl

import (
	"ForumPublica/server/services"
	"ForumPublica/server/tasks"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//AppMarketItems /app/market/index
func AppMarketItems(c *gin.Context) {
	u := user(c)
	p := page(c)

	c.Keys["MarketItems"] = services.MarketItemsList(u.ID, p)
	c.Keys["Title"] = "Market Items"
	c.Keys["chars"] = services.CharsByUserID(u.ID)

	c.HTML(http.StatusOK, "app/market_items/index.html", c.Keys)
}

//AppMarketItemsAdd add market item
func AppMarketItemsAdd(c *gin.Context) {
	u := user(c)

	params := make(map[string]int32)
	c.BindJSON(&params)

	services.MarketItemsCreate(u.ID, params)

	p := page(c)

	list := services.MarketItemsList(u.ID, p)

	c.JSON(http.StatusOK, list)
}

//AppMarketItemsDelete delete market item
func AppMarketItemsDelete(c *gin.Context) {
	u := user(c)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	services.MarketItemsDelete(u.ID, id)

	c.Redirect(http.StatusTemporaryRedirect, "/app/market_items")
}

//AppMarketItemsSyncQty sync qty
func AppMarketItemsSyncQty(c *gin.Context) {
	u := user(c)
	tasks.LoadMarketDataForUser(u.ID)
	c.Redirect(http.StatusTemporaryRedirect, "/app/market_items")
}
