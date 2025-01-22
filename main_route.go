package main

import (
	"Blog-CMS/component/appctx"
	"Blog-CMS/middleware"
	usertransport "Blog-CMS/module/user/transport"
	"github.com/gin-gonic/gin"
)

func SetupGroup(appCtx appctx.AppContext, v1 *gin.RouterGroup) {

	v1.POST("/register", usertransport.Register(appCtx))
	v1.POST("/authenticate", usertransport.Login(appCtx))
	v1.GET("/profile", middleware.RequireAuth(appCtx), usertransport.Profile(appCtx))
	v1.GET("/user/paging", middleware.RequireAuth(appCtx), usertransport.ListUser(appCtx))
}
