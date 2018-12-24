package ctrl

import (
	"fmt"
	"ForumPublica/server/middleware"
	"ForumPublica/server/models"
	"ForumPublica/server/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//AppConstructionExpensesAdd add expense
func AppConstructionExpensesAdd(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	cid, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	params := services.ExpenseParams{}
	err := c.BindJSON(&params)
	if err!=nil {
		fmt.Println("AppConstructionExpensesAdd",err)
	}

	services.ConstructionExpenseAdd(user.ID, cid, params)
	cn, _ := services.ConstructionGet(user.ID, cid)

	c.JSON(http.StatusOK, cn)
}

func AppConstructionExpensesDelete(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	cid, _ := strconv.ParseInt(c.Param("cid"), 10, 64)

	services.ConstructionExpenseDelete(user.ID, cid, id)
	cn, _ := services.ConstructionGet(user.ID, cid)

	c.JSON(http.StatusOK, cn)
}
