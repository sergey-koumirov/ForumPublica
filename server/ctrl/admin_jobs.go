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
