package ctrl

import (
	"ForumPublica/server/middleware"
	"ForumPublica/server/models"
	"ForumPublica/server/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AppConstructionsAddRun(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	params := make(map[string]int64)
	c.BindJSON(&params)

	services.ConstructionRunAdd(user.Id, id, params)
	cn, _ := services.ConstructionGet(user.Id, id)

	c.JSON(http.StatusOK, cn)
}

func AppConstructionsDeleteRun(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	cid, _ := strconv.ParseInt(c.Param("cid"), 10, 64)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	services.ConstructionRunDelete(user.Id, cid, id)
	cn, _ := services.ConstructionGet(user.Id, cid)

	c.JSON(http.StatusOK, cn)
}

