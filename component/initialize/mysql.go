package initialize

import (
	"Blog-CMS/common"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitMysql() *gorm.DB {
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

	return db
}