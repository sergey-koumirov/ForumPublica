package ctrl

import (
	"ForumPublica/server/middleware"
	"ForumPublica/server/models"
	"ForumPublica/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AppSAP(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	chars := services.CharJobsList(user.Id)

	c.Keys["chars"] = chars
	c.Keys["timeout"] = services.GetTimeout(services.JOBS, 5)
	c.HTML(http.StatusOK, "app/sap/index.html", c.Keys)
}

func AppSAPRefresh(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	services.RefreshJobs(user.Id)

	c.Redirect(http.StatusTemporaryRedirect, "/app/sap")
}
