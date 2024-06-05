package models

import "GameManageSystem/models/ctype"

type UserModel struct {
	ID       string     `gorm:"size:20;comment:学号" json:"id"`
	Name     string     `gorm:"size:32;comment:姓名" json:"name"`
	Password string     `gorm:"size:128;comment:密码" json:"password"`
	Academy  string     `gorm:"size:128;comment:学院" json:"academy"`
	Gender   string     `gorm:"size:16;comment:性别" json:"gender"`
	Major    string     `gorm:"size:128;comment:专业" json:"major"`
	Role     ctype.Role `gorm:"size:4;default:2;comment:权限：1管理员,2学生,3老师" json:"role,select(info)"` // 权限  1 管理员  2 学生  3 老师
}
