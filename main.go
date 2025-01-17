package main

import (
	"Blog-CMS/component/initialize"
	"Blog-CMS/middleware"
)

func main() {

	// load config init
	r, appCtx := initialize.RunInit()
	r.Use(middleware.Recover(appCtx))

	blog := r.Group("/blog")

	SetupGroup(appCtx, blog)

	r.Run() // listen and serve on 0.0.0.0:8080
}
