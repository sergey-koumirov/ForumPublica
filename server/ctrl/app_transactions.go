package ctrl

import (
	"ForumPublica/server/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

//AppMarketItems /app/market/index
func AppTransactions(c *gin.Context) {
	u := user(c)
	p := page(c)

	c.Keys["Transactions"] = services.TransactionsList(u.ID, p)
	c.Keys["Title"] = "Transactions"

	c.HTML(http.StatusOK, "app/transactions/index.html", c.Keys)
}
