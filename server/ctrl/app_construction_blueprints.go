package ctrl

import (
	"ForumPublica/server/middleware"
	"ForumPublica/server/models"
	"ForumPublica/server/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//AppConstructionsAddBlueprint add
func AppConstructionsAddBlueprint(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	params := make(map[string]int32)
	c.BindJSON(&params)

	services.ConstructionBluprintAdd(user.ID, id, params)
	cn, _ := services.ConstructionGet(user.ID, id)

	c.JSON(http.StatusOK, cn)
}

//AppConstructionBlueprintsDelete delete
func AppConstructionBlueprintsDelete(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	cid, _ := strconv.ParseInt(c.Param("cid"), 10, 64)

	services.ConstructionBlueprintDelete(user.ID, cid, id)
	cn, _ := services.ConstructionGet(user.ID, cid)

	c.JSON(http.StatusOK, cn)
}

//AppConstructionBlueprintsUpdate update
func AppConstructionBlueprintsUpdate(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	cid, _ := strconv.ParseInt(c.Param("cid"), 10, 64)

	params := make(map[string]int32)
	c.BindJSON(&params)

	services.ConstructionBlueprintUpdate(user.ID, cid, id, params)
	cn, _ := services.ConstructionGet(user.ID, cid)

	c.JSON(http.StatusOK, cn)
}
