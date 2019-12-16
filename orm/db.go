package orm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var GlobalDb *gorm.DB

func init() {
	/**
	db, err := gorm.Open("mysql", "cron:admin@(10.70.30.26:3306)/cron?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	GlobalDb = db
	**/
}
