package ctrl

import (
	"net/http"

	"ForumPublica/server/tasks"

	"github.com/gin-gonic/gin"
)

//AdminJobs page with buttons to start any async job
func AdminJobs(c *gin.Context) {
	c.Keys["Title"] = "Admin: Jobs"
	c.HTML(http.StatusOK, "admin/jobs.html", c.Keys)
}

//AdminJobsUpdatePrices start update tasks job
func AdminJobsCheckT2(c *gin.Context) {
	tasks.TaskCheckT2()
	c.Redirect(http.StatusTemporaryRedirect, "/admin/jobs")
}

//AdminJobsTestMarket start test market job
func AdminJobsLoadTransactions(c *gin.Context) {
	tasks.TaskLoadTransactions()
	c.Redirect(http.StatusTemporaryRedirect, "/admin/jobs")
}

//AdminJobsLoadMarketData load market data
func AdminJobsLoadMarketData(c *gin.Context) {
	tasks.LoadMarketData()
	c.Redirect(http.StatusTemporaryRedirect, "/admin/jobs")
}
