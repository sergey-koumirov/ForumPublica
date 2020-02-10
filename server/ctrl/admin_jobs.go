package ctrl

import (
	"net/http"

	"ForumPublica/server/tasks"

	"github.com/gin-gonic/gin"
)

//AdminJobs page with buttons to start any async job
func AdminJobs(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/jobs.html", c.Keys)
}

//AdminJobsUpdatePrices start update tasks job
func AdminJobsUpdatePrices(c *gin.Context) {

	tasks.TaskUpdatePrices()

	c.Redirect(http.StatusTemporaryRedirect, "/admin/jobs")
}

//AdminJobsTestMarket start test market job
func AdminJobsTestMarket(c *gin.Context) {
	u := user(c)
	tasks.TaskTestMarket(u)
	c.Redirect(http.StatusTemporaryRedirect, "/admin/jobs")
}

//AdminJobsLoadMarketData load market data
func AdminJobsLoadMarketData(c *gin.Context) {
	tasks.LoadMarketData()
	c.Redirect(http.StatusTemporaryRedirect, "/admin/jobs")
}
