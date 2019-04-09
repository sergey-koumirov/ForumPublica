package ctrl

import (
	"ForumPublica/server/services"
	"ForumPublica/server/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//AppConstructions list
func AppConstructions(c *gin.Context) {
	u := user(c)

	p := page(c)
	list := services.ConstructionsList(u.ID, p)

	c.Keys["constructions"] = list
	c.Keys["p"] = utils.NewPagination(list.Total, services.PerPage, p, "/app/constructions")

	c.HTML(http.StatusOK, "app/constructions/index.html", c.Keys)
}

//AppConstructionsAdd add
func AppConstructionsAdd(c *gin.Context) {
	u := user(c)

	new := services.ConstructionCreate(u.ID)
	c.Keys["construction"] = new

	c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("/app/construction/%d", new.ID))
}

//AppConstructionsShow show
func AppConstructionsShow(c *gin.Context) {
	u := user(c)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	cn, err := services.ConstructionGet(u.ID, id)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	c.Keys["construction"] = cn
	c.Keys["chars"] = services.CharsByUserID(u.ID)

	c.HTML(http.StatusOK, "app/constructions/show.html", c.Keys)
}

//AppConstructionsDelete delete
func AppConstructionsDelete(c *gin.Context) {
	u := user(c)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	services.ConstructionDelete(u.ID, id)

	c.Redirect(http.StatusTemporaryRedirect, "/app/constructions")
}

//AppConstructionsSaveBonus save bonus
func AppConstructionsSaveBonus(c *gin.Context) {
	u := user(c)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	params := make(map[string]string)
	c.BindJSON(&params)

	services.ConstructionSaveBonus(u.ID, id, params)

	cn, _ := services.ConstructionGet(u.ID, id)
	c.JSON(http.StatusOK, cn)
}
