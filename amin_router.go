package main

import (
	"Blog-CMS/component/appctx"
	"Blog-CMS/middleware"
	usertransport "Blog-CMS/module/user/transport"
	"github.com/gin-gonic/gin"
)

func SetupAdminRoute(appCtx appctx.AppContext, v1 *gin.RouterGroup) {
	admin := v1.Group("/admin",
		middleware.RequireAuth(appCtx),
		middleware.RoleRequired(appCtx, "admin", "mod"),
	)
	{
		admin.GET("/profile", middleware.RequireAuth(appCtx), usertransport.Profile(appCtx))
	}
}
