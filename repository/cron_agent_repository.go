package repository

import (
	"crontab/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var CronAgent = cronAgentRepository{}

type cronAgentRepository struct {
}

func (t *cronAgentRepository) Save(agent entity.CronAgent) {
	db, err := gorm.Open("mysql", "root:admin@(10.70.30.26:3306)/cron?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	defer func() {
		_ = db.Close()
	}()
	db.Create(agent)
}
