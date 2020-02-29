package ctrl

import (
	"ForumPublica/server/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

//AppMarketItems /app/market/index
func AppSummary(c *gin.Context) {
	u := user(c)

	c.Keys["Summary"] = services.TransactionsSummary(u.ID)
	c.Keys["Title"] = "Summary"

	c.HTML(http.StatusOK, "app/summary/index.html", c.Keys)
}
