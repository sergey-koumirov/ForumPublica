package middleware

import (
	"ForumPublica/server/esi"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {

	value, _ := c.Get("user")

	if value == nil {
		url, state := esi.CallbackURL()

		session := sessions.Default(c)
		session.Set("state", state)
		session.Save()

		c.Redirect(http.StatusTemporaryRedirect, url)
		c.Abort()
		return
	}

}

func AuthCallback(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("state")
	if value == nil {
		return
	}
	saved := value.(string)

	fmt.Println(saved)
	fmt.Printf("%+v\n", c)

}
