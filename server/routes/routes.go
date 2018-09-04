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
		authorized.POST("/construction/:id/save_bonus", ctrl.AppConstructionsSaveBonus)

		authorized.GET("/search/item_type", ctrl.AppSearchItemType)
	}
}
