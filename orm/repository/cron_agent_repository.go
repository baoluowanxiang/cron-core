package repository

import (
	"crontab/orm/entity"
	"github.com/jinzhu/gorm"
)

type CronAgentRepository struct {

}

func (t *CronAgentRepository) Save(agent entity.CronAgent) {
	db, err := gorm.Open("mysql", "root:admin@/cron?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer func() {
		_ = db.Close()
	}()
	db.Create(agent)
}

