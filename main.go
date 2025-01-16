package main

import (
	"Blog-CMS/common"
	"Blog-CMS/component/appctx"
	"Blog-CMS/component/initialize"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {

	// load config init
	r := initialize.RunInit()

	// init db
	m := common.Config.MySql

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.UserName, m.Password, m.Host, m.Port, m.Dbname)

	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})

	if err != nil {
		log.Fatalln(db, err)
	}

	appCtx := appctx.NewAppContext(db)

	blog := r.Group("/blog")

	SetupGroup(appCtx, blog)

	r.Run() // listen and serve on 0.0.0.0:8080
}
