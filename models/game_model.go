package models

type GameModel struct {
	Name      string `gorm:"size32;comment:比赛名称 primary_key" json:"name" `
	Organizer string `gorm:"size128;comment:主办方" json:"organizer"`
	StartTime string `gorm:"size256;comment:比赛开始时间" json:"start_time"`
	EndTime   string `gorm:"size256;comment:比赛结束时间" json:"end_time"`
	Address   string `gorm:"size128;comment:比赛地点" json:"address"`
	Prize     string `gorm:"size128;comment:比赛奖品" json:"prize"`
	Creator   string `gorm:"size128;comment:创建者" json:"creator"`
}
