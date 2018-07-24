package ctrl

import (
	"ForumPublica/server/db"
	"ForumPublica/server/middleware"
	"ForumPublica/server/models"
	"ForumPublica/server/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func AppSAP(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	var chars []models.Character

	errSel := db.DB.Preload("Jobs", jobsOrder).Where("user_id=?", user.Id).Order("name").Find(&chars).Error

	if errSel != nil {
		fmt.Println("errSel", errSel)
	}

	c.Keys["chars"] = chars

	c.HTML(http.StatusOK, "app/sap.html", c.Keys)
}

func jobsOrder(db *gorm.DB) *gorm.DB {
	return db.Order("esi_jobs.end_date desc")
}

func AppSAPRefresh(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	services.RefreshJobs(user.Id)

	c.Redirect(http.StatusTemporaryRedirect, "/app/sap")
}
