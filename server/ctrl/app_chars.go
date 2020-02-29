package ctrl

import (
	"ForumPublica/server/db"
	"ForumPublica/server/models"
	"ForumPublica/server/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//AppChars list
func AppChars(c *gin.Context) {
	u := user(c)

	var chars []models.Character
	db.DB.Preload("Skills", skillsOrder).Where("user_id = ?", u.ID).Order("name").Find(&chars)

	c.Keys["chars"] = chars
	c.Keys["Title"] = "Characters"

	c.HTML(http.StatusOK, "app/chars.html", c.Keys)
}

func skillsOrder(db *gorm.DB) *gorm.DB {
	return db.Order("esi_skills.name asc")
}

//CharRefreshSkills refresh skills
func CharRefreshSkills(c *gin.Context) {
	u := user(c)

	cid, _ := strconv.ParseInt(c.Param("cid"), 10, 64)
	char := models.Character{}
	errSel := db.DB.Where("id = ? and user_id = ?", cid, u.ID).First(&char).Error

	if errSel == nil {
		services.RefreshSkills(cid)
	} else {
		fmt.Println("errSel", errSel)
	}

	c.Redirect(http.StatusTemporaryRedirect, "/app/chars")
}
