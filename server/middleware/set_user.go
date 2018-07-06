package middleware

import (
	"ForumPublica/server/db"
	"ForumPublica/server/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetUser(c *gin.Context) {

	session := sessions.Default(c)
	value := session.Get("userId")
	if value == nil {
		return
	}

	var (
		userId int64 = value.(int64)
		user         = models.User{Id: userId}
	)
	ex, errDb := db.DB.Get(&user)
	if !ex || errDb != nil {
		return
	}

	c.Set("user", user)

}
