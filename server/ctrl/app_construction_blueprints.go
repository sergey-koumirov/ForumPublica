package ctrl

import (
	"ForumPublica/server/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//AppConstructionsAddBlueprint add
func AppConstructionsAddBlueprint(c *gin.Context) {
	u := user(c)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	params := make(map[string]int32)
	c.BindJSON(&params)

	services.ConstructionBluprintAdd(u.ID, id, params)
	cn, _ := services.ConstructionGet(u.ID, id)

	c.JSON(http.StatusOK, cn)
}

//AppConstructionBlueprintsDelete delete
func AppConstructionBlueprintsDelete(c *gin.Context) {
	u := user(c)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	cid, _ := strconv.ParseInt(c.Param("cid"), 10, 64)

	services.ConstructionBlueprintDelete(u.ID, cid, id)
	cn, _ := services.ConstructionGet(u.ID, cid)

	c.JSON(http.StatusOK, cn)
}

//AppConstructionBlueprintsUpdate update
func AppConstructionBlueprintsUpdate(c *gin.Context) {
	u := user(c)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	cid, _ := strconv.ParseInt(c.Param("cid"), 10, 64)

	params := make(map[string]int32)
	c.BindJSON(&params)

	services.ConstructionBlueprintUpdate(u.ID, cid, id, params)
	cn, _ := services.ConstructionGet(u.ID, cid)

	c.JSON(http.StatusOK, cn)
}
