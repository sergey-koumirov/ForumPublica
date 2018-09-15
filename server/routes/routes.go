package routes

import (
	"ForumPublica/server/ctrl"
	"ForumPublica/server/middleware"

	"github.com/gin-gonic/gin"
)

func AddAppRoutes(r *gin.Engine) {
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
		authorized.GET("/constructions/add", ctrl.AppConstructionsAdd)
		authorized.GET("/construction/:id", ctrl.AppConstructionsShow)
		authorized.GET("/construction/:id/delete", ctrl.AppConstructionsDelete)
		authorized.POST("/construction/:id/save_bonus", ctrl.AppConstructionsSaveBonus)
		authorized.POST("/construction/:id/add_blueprint", ctrl.AppConstructionsAddBlueprint)

		authorized.DELETE("/construction/:cid/blueprint/:id", ctrl.AppConstructionBlueprintsDelete)
		authorized.PATCH("/construction/:cid/blueprint/:id", ctrl.AppConstructionBlueprintsUpdate)

		authorized.GET("/search/:filter", ctrl.AppSearchItemType)

		authorized.POST("/ui/open_market", ctrl.AppUIOpenMarket)

	}

	admin := r.Group("/admin")
	admin.Use(middleware.Auth, middleware.Admin)
	{
		admin.GET("/index", ctrl.AdminIndex)
		admin.GET("/users", ctrl.AdminUsers)
	}

}
