package ctrl

import (
	"ForumPublica/server/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//AppConstructionsAddRun add run
func AppConstructionsAddRun(c *gin.Context) {
	u := user(c)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	params := make(map[string]int64)
	c.BindJSON(&params)

	services.ConstructionRunAdd(u.ID, id, params)
	cn, _ := services.ConstructionGet(u.ID, id)

	c.JSON(http.StatusOK, cn)
}

//AppConstructionsDeleteRun delete run
func AppConstructionsDeleteRun(c *gin.Context) {
	u := user(c)

	cid, _ := strconv.ParseInt(c.Param("cid"), 10, 64)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	services.ConstructionRunDelete(u.ID, cid, id)
	cn, _ := services.ConstructionGet(u.ID, cid)

	c.JSON(http.StatusOK, cn)
}
