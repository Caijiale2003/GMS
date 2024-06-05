package models

type EnrollModel struct {
	GameName string `gorm:"size32;comment:比赛名称" json:"game_name"`
	ID       string `gorm:"size20;comment:学号" json:"id"`
	Name     string `gorm:"size32;comment:姓名" json:"name"`
	Teacher  string `gorm:"size32;comment:指导老师" json:"teacher"`
	TeamName string `gorm:"size32;comment:队伍名称" json:"team_name"`
}
