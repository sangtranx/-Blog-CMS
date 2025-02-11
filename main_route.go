package main

import (
	"Blog-CMS/component/appctx"
	"Blog-CMS/middleware"
	posttransport "Blog-CMS/module/post/transport"
	postliketranspot "Blog-CMS/module/postlike/transport"
	usertransport "Blog-CMS/module/user/transport"

	"github.com/gin-gonic/gin"
)

func SetupGroup(appCtx appctx.AppContext, v1 *gin.RouterGroup) {

	// user
	v1.POST("/register", usertransport.Register(appCtx))
	v1.POST("/login", usertransport.Login(appCtx))
	v1.POST("/logout", middleware.RequireAuth(appCtx), usertransport.Logout(appCtx))
	v1.POST("/change/password", middleware.RequireAuth(appCtx), usertransport.UpdatePassword(appCtx))
	v1.GET("/profile", middleware.RequireAuth(appCtx), usertransport.Profile(appCtx))
	v1.GET("/user/paging", middleware.RequireAuth(appCtx), usertransport.ListUser(appCtx))

	// post
	v1.POST("/post/create", middleware.RequireAuth(appCtx), posttransport.CreateNewPost(appCtx))
	v1.POST("/post/delete", middleware.RequireAuth(appCtx), posttransport.DeletePost(appCtx))
	v1.POST("/post/update", middleware.RequireAuth(appCtx), posttransport.UpdatePost(appCtx))

	v1.POST("/post/like", middleware.RequireAuth(appCtx), postliketranspot.UserlikePost(appCtx))
	v1.DELETE("/post/dislike", middleware.RequireAuth(appCtx), postliketranspot.UserDislikePost(appCtx))

}
