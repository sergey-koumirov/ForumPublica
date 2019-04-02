package ctrl

import (
	"ForumPublica/server/middleware"
	"ForumPublica/server/models"
	"ForumPublica/server/services"
	"ForumPublica/server/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//AppMarketItems /app/market/index
func AppMarketItems(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	page, errParse := strconv.ParseInt(c.Query("page"), 10, 64)
	if errParse != nil {
		page = 1
	}
	list := services.MarketItemsList(user.ID, page)

	c.Keys["MarketItems"] = list
	c.Keys["p"] = utils.NewPagination(list.Total, services.PerPage, page, "/app/market_items")

	c.HTML(http.StatusOK, "app/market_items/index.html", c.Keys)
}

//AppMarketItemsAdd add market item
func AppMarketItemsAdd(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	new := services.MarketItemsCreate(user.ID)
	c.Keys["MarketItem"] = new

	c.Redirect(http.StatusTemporaryRedirect, "/app/market_items")
}

//AppMarketItemsDelete delete market item
func AppMarketItemsDelete(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	services.MarketItemsDelete(user.ID, id)

	c.Redirect(http.StatusTemporaryRedirect, "/app/market_items")
}