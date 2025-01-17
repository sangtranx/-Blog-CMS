package main

import (
	"Blog-CMS/component/initialize"
)

func main() {

	// load config init
	r, appCtx := initialize.RunInit()

	blog := r.Group("/blog")

	SetupGroup(appCtx, blog)

	r.Run() // listen and serve on 0.0.0.0:8080
}
