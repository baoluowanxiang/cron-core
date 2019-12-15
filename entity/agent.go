package entity

type CronAgent struct {
	Id      uint   `gorm:"primary_key"`
	Service string `gorm:"service"`
	Ip      string `gorm:"ip"`
	Status  int    `gorm:"status"`
}

const (
	CronAgentOnLine  = 1
	CronAgentOffLine = 2
)
