package ctrl

import (
	"ForumPublica/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AppSAP index
func AppSAP(c *gin.Context) {
	u := user(c)
	chars := services.CharJobsList(u.ID)
	c.Keys["chars"] = chars
	c.Keys["timeout"] = services.GetTimeout(services.JOBS, 5)
	c.HTML(http.StatusOK, "app/sap/index.html", c.Keys)
}

//AppSAPRefresh refresh
func AppSAPRefresh(c *gin.Context) {
	u := user(c)
	services.RefreshJobs(u.ID)
	c.Redirect(http.StatusTemporaryRedirect, "/app/sap")
}
