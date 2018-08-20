package main

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/config"
	"ForumPublica/server/ctrl"
	"ForumPublica/server/db"
	"ForumPublica/server/middleware"
	"ForumPublica/server/utils"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

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

	static.LoadJSONs(config.Vars.SDE)

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
	defer db.DB.Close()

	store := cookie.NewStore([]byte(config.Vars.SESSION_KEY))

	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"TimeoutClass": utils.TimeoutClass,
	})

	r.Static("/assets", "./server/assets")
	r.StaticFile("/favicon.ico", "./server/assets/favicon.ico")

	r.Delims("<%", "%>")

	templates := make([]string, 0)
	filepath.Walk("./server/templates", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			templates = append(templates, path)
		}
		return nil
	})
	r.LoadHTMLFiles(templates...)

	r.Use(sessions.Sessions("mysession", store))
	r.Use(middleware.SetUser)

	r.GET("/", ctrl.RootIndex)
	r.GET("/login", middleware.Auth)
	r.GET("/logout", middleware.Auth, middleware.Logout)
	r.GET("/probleme_callback", middleware.AuthCallback)

	authorized := r.Group("/app")
	authorized.Use(middleware.Auth)
	{
		authorized.GET("/index", ctrl.AppIndex)

		authorized.GET("/chars", ctrl.AppChars)
		authorized.GET("/chars/add", middleware.Add)
		authorized.GET("/char/:cid/refresh_skills", ctrl.CharRefreshSkills)

		authorized.GET("/sap", ctrl.AppSAP)
		authorized.GET("/sap/refresh", ctrl.AppSAPRefresh)

		authorized.GET("/constructions", ctrl.AppConstructions)
	}

	gin.SetMode(config.Vars.MODE)
	r.Run(":" + config.Vars.PORT)
}
