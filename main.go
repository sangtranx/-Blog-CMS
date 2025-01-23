package main

import (
	_ "Blog-CMS/cmd/swag/docs"
	"Blog-CMS/component/initialize"
	"Blog-CMS/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os/exec"
	"runtime"
)

// @title           API Document Blog-CMS
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  https://github.com/sangtranx/-Blog-CMS/tree/develop

// @contact.name   API Support
// @contact.url    https://github.com/sangtranx/-Blog-CMS/tree/develop
// @contact.email  tran.thanhsang.dev@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /blog/
// @schema http
func main() {

	// load config init
	r, appCtx := initialize.RunInit()
	r.Use(middleware.Recover(appCtx))

	blog := r.Group("/blog")

	SetupAdminRoute(appCtx, blog)
	SetupGroup(appCtx, blog)

	// Generate swagger docs
	generateSwaggerDocs()
	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Open Swagger UI in browser
	//openBrowser("http://localhost:8080/swagger/index.html")

	r.Run() // listen and serve on 0.0.0.0:8080

}

func generateSwaggerDocs() {
	cmd := exec.Command("swag", "init", "-g", "main.go", "-o", "cmd/swag/docs")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to generate Swagger docs: %s\n%s", err, output)
	}
	log.Println("Swagger docs generated successfully in cmd/swag/docs")
}

// openBrowser opens the specified URL in the default browser
func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = exec.Command("xdg-open", url).Start()
	}

	if err != nil {
		log.Printf("Failed to open browser: %v\n", err)
	}
}
