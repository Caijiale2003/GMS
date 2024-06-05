package models

type ScoreModel struct {
	GameName string `gorm:"size:32;comment:比赛名称" json:"game_name"`
	ID       string `gorm:"size:20;comment:学号" json:"id"`
	Score    int    `gorm:"size:16;comment:分数" json:"score"`
}
