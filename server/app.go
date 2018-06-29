package main

import(
	"github.com/gin-gonic/gin"
	"ForumPublica/server/config"
	"ForumPublica/server/middleware"
	"ForumPublica/server/ctrl"
	"fmt"
)


func main() {

	config.LoadVars()
	if config.Vars == nil {
		fmt.Println("Vars load problem")
		return
	}

	r := gin.Default()

  r.GET("/login", ctrl.Login)


	authorized := r.Group("/app")
	authorized.Use(middleware.Auth())
	{
		authorized.GET("/index", ctrl.Index)
	}

  gin.SetMode(config.Vars.MODE)
	r.Run(":"+config.Vars.PORT)
}
