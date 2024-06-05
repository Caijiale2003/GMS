package models

type TeamModel struct {
	TeamName string `gorm:"size:32;comment:队伍名称" json:"team_name"`
	ID       string `gorm:"size:20;comment:学号" json:"id"`
	Name     string `gorm:"size:32;comment:姓名" json:"name"`
}
