package main

import (
	"Blog-CMS/component/appctx"
	"Blog-CMS/middleware"
	chattransport "Blog-CMS/module/chat/transport"
	"github.com/gin-gonic/gin"
)

func SetupChatAi(appCtx appctx.AppContext, v1 *gin.RouterGroup) {
	// deepseek
	v1.POST("/deepseek/chat", middleware.RequireAuth(appCtx), chattransport.HandleDeepSeekQuery(appCtx))
}
