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

//AppMarket /app/market/index
func AppMarket(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	page, errParse := strconv.ParseInt(c.Query("page"), 10, 64)
	if errParse != nil {
		page = 1
	}
	list := services.MarketItemsList(user.ID, page)

	c.Keys["MarketItems"] = list
	c.Keys["p"] = utils.NewPagination(list.Total, services.PerPage, page, "/app/market")

	c.HTML(http.StatusOK, "app/market/index.html", c.Keys)
}
