package main

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/config"
	"ForumPublica/server/db"
	"ForumPublica/server/middleware"
	"ForumPublica/server/routes"
	"ForumPublica/server/services"
	"ForumPublica/server/tasks"
	"ForumPublica/server/utils"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/astaxie/beego/toolbox"
)

func main() {

	config.LoadVars()
	if config.Vars == nil {
		fmt.Println("Vars load problem")
		return
	}

	static.LoadJSONs(config.Vars.SDE)
	runtime.GC()

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

	services.InitPrices()

	store := cookie.NewStore([]byte(config.Vars.SESSION_KEY))

	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"TimeoutClass": utils.TimeoutClass,
		"Marshal":      utils.Marshal,
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

	routes.AddAppRoutes(r)

	toolbox.AddTask("update_prices", toolbox.NewTask("update_prices", "0 22 14 * * *", tasks.TaskUpdatePrices))
	toolbox.StartTask()
	defer toolbox.StopTask()

	gin.SetMode(config.Vars.MODE)
	r.Run(":" + config.Vars.PORT)
}
