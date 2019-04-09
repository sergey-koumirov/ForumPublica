package ctrl

import (
	"ForumPublica/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AppUIOpenMarket open market
func AppUIOpenMarket(c *gin.Context) {
	u := user(c)

	params := make(map[string]int64)
	c.BindJSON(&params)

	services.UIOpenMarket(u.ID, params)

	c.JSON(http.StatusOK, gin.H{})
}
