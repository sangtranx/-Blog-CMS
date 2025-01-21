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
		admin.GET("/profile", usertransport.Profile(appCtx))
		admin.GET("/userProfile", usertransport.GetUserProfile(appCtx))
		admin.GET("/users", usertransport.GetAllUser(appCtx))
	}
}
