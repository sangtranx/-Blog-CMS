package initialize

import (
	"Blog-CMS/common"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitMysql() {

	m := common.Config.MySql

	dsn := "%s:%s@tpc(%v:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	s := fmt.Sprintf(dsn, m.UserName, m.Password, m.Host, m.Port, m.Dbname)

	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})

	if err != nil {
		log.Fatalln(db, err)
	}

	common.DB = db
}
