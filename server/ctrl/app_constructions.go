package ctrl

import (
	"ForumPublica/server/middleware"
	"ForumPublica/server/models"
	"ForumPublica/server/services"
	"ForumPublica/server/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AppConstructions(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	page, errParse := strconv.ParseInt(c.Query("page"), 10, 64)
	if errParse != nil {
		page = 1
	}
	list := services.ConstructionsList(user.ID, page)

	c.Keys["constructions"] = list
	c.Keys["p"] = utils.NewPagination(list.Total, services.PerPage, page, "/app/constructions")

	c.HTML(http.StatusOK, "app/constructions/index.html", c.Keys)
}

func AppConstructionsAdd(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	new := services.ConstructionCreate(user.ID)
	c.Keys["construction"] = new

	c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("/app/construction/%d", new.ID))
}

func AppConstructionsShow(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	cn, err := services.ConstructionGet(user.ID, id)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	c.Keys["construction"] = cn
	c.Keys["chars"] = services.CharsByUserID(user.ID)

	c.HTML(http.StatusOK, "app/constructions/show.html", c.Keys)
}

func AppConstructionsDelete(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	services.ConstructionDelete(user.ID, id)

	c.Redirect(http.StatusTemporaryRedirect, "/app/constructions")
}

func AppConstructionsSaveBonus(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	params := make(map[string]string)
	c.BindJSON(&params)

	services.ConstructionSaveBonus(user.ID, id, params)

	cn, _ := services.ConstructionGet(user.ID, id)
	c.JSON(http.StatusOK, cn)
}
