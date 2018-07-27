package ctrl

import (
	"ForumPublica/server/db"
	"ForumPublica/server/middleware"
	"ForumPublica/server/models"
	"ForumPublica/server/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func AppChars(c *gin.Context) {

	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	var chars []models.Character
	db.DB.Preload("Skills", skillsOrder).Where("user_id = ?", user.Id).Order("name").Find(&chars)

	c.Keys["chars"] = chars

	c.HTML(http.StatusOK, "app/chars.html", c.Keys)
}

func skillsOrder(db *gorm.DB) *gorm.DB {
	return db.Order("esi_skills.name asc")
}

func CharRefreshSkills(c *gin.Context) {
	raw, _ := c.Get(middleware.USER)
	user := raw.(models.User)

	cid, _ := strconv.ParseInt(c.Param("cid"), 10, 64)
	char := models.Character{}
	errSel := db.DB.Where("id = ? and user_id = ?", cid, user.Id).First(&char).Error

	if errSel == nil {
		services.RefreshSkills(cid)
	} else {
		fmt.Println("errSel", errSel)
	}

	c.Redirect(http.StatusTemporaryRedirect, "/app/chars")
}
