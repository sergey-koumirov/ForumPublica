package ctrl

import (
	"ForumPublica/server/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//AppConstructionExpensesAdd add expense
func AppConstructionExpensesAdd(c *gin.Context) {
	u := user(c)

	cid, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	params := services.ExpenseParams{}
	err := c.BindJSON(&params)
	if err != nil {
		fmt.Println("AppConstructionExpensesAdd", err)
	}

	services.ConstructionExpenseAdd(u.ID, cid, params)
	cn, _ := services.ConstructionGet(u.ID, cid)

	c.JSON(http.StatusOK, cn)
}

//AppConstructionExpensesDelete delete expense
func AppConstructionExpensesDelete(c *gin.Context) {
	u := user(c)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	cid, _ := strconv.ParseInt(c.Param("cid"), 10, 64)

	services.ConstructionExpenseDelete(u.ID, cid, id)
	cn, _ := services.ConstructionGet(u.ID, cid)

	c.JSON(http.StatusOK, cn)
}
