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
	list := services.ConstructionsList(user.Id, page)

	c.Keys["constructions"] = list
	c.Keys["p"] = utils.NewPagination(list.Total, services.PER_PAGE, page, "/app/constructions")

	c.HTML(http.StatusOK, "app/constructions/index.html", c.Keys)
}

func AppConstructionsAdd(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	new := services.ConstructionCreate(user.Id)
	c.Keys["construction"] = new

	c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("/app/construction/%d", new.Id))
}

func AppConstructionsShow(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	cn, err := services.ConstructionGet(user.Id, id)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	c.Keys["construction"] = cn

	c.HTML(http.StatusOK, "app/constructions/show.html", c.Keys)
}

func AppConstructionsSaveBonus(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	params := make(map[string]string)
	c.BindJSON(&params)

	services.ConstructionSaveBonus(user.Id, id, params)
	c.JSON(http.StatusOK, gin.H{})
}
