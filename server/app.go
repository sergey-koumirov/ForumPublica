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
	"flag"
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

	resetCache := flag.Bool("reset-cache", false, "true/false")
	flag.Parse()
	static.LoadJSONs(config.Vars.SDE, *resetCache)
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

	store := cookie.NewStore([]byte(config.Vars.SessionKey))

	gin.SetMode(config.Vars.MODE)
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"TimeoutClass": utils.TimeoutClass,
		"Marshal":      utils.Marshal,
		"FormatFloat":  utils.FormatFloat,
		"FormatInt":    utils.FormatInt,
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

	toolbox.AddTask("load_market_data", toolbox.NewTask("load_market_data", "0 30 */4 * * *", tasks.LoadMarketData))
	toolbox.AddTask("load_transactions", toolbox.NewTask("load_transactions", "0 32 */2 * * *", tasks.TaskLoadTransactions))
	toolbox.AddTask("load_deviations", toolbox.NewTask("load_deviations", "0 34 0 * * *", tasks.TaskCheckT2))
	toolbox.StartTask()
	defer toolbox.StopTask()

	r.RunTLS(":"+config.Vars.PORT, config.Vars.SSLPath+"nginx.crt", config.Vars.SSLPath+"nginx.key")
}
