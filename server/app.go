package main

import (
	"ForumPublica/server/config"
	"ForumPublica/server/ctrl"
	"ForumPublica/server/db"
	"ForumPublica/server/middleware"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadVars()
	if config.Vars == nil {
		fmt.Println("Vars load problem")
		return
	}

	errMgr := db.Migrate()
	if errMgr != nil {
		fmt.Println("Migration problems: ", errMgr)
		return
	}

	db.Connect()
	if db.DB == nil {
		fmt.Println("No database connection")
		return
	}

	store := cookie.NewStore([]byte(config.Vars.SESSION_KEY))

	r := gin.Default()

	r.Static("/assets", "./server/assets")
	r.StaticFile("/favicon.ico", "./server/assets/favicon.ico")

	r.LoadHTMLGlob("server/templates/**/*.html")

	r.Use(sessions.Sessions("mysession", store))
	r.Use(middleware.SetUser)

	r.GET("/", ctrl.RootIndex)
	r.GET("/probleme_callback", middleware.AuthCallback)

	authorized := r.Group("/app")
	authorized.Use(middleware.Auth)
	{
		authorized.GET("/index", ctrl.AppIndex)
		authorized.GET("/chars", ctrl.AppChars)
	}

	gin.SetMode(config.Vars.MODE)
	r.Run(":" + config.Vars.PORT)
}
