package ctrl

import (
	"ForumPublica/server/db"
	"ForumPublica/server/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func AdminUsers(c *gin.Context) {

	var users []models.User
	db.DB.Preload("Characters", charsOrder).Order("id").Find(&users)

	c.Keys["users"] = users

	c.HTML(http.StatusOK, "admin/users.html", c.Keys)
}

func charsOrder(db *gorm.DB) *gorm.DB {
	return db.Order("esi_characters.name asc")
}
