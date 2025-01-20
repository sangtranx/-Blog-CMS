package main

import (
	"Blog-CMS/component/appctx"
	usertransport "Blog-CMS/module/user/transport"
	"github.com/gin-gonic/gin"
)

func SetupGroup(appCtx appctx.AppContext, v1 *gin.RouterGroup) {

	v1.POST("/register", usertransport.Register(appCtx))
	v1.POST("/authenticate", usertransport.Login(appCtx))
}
