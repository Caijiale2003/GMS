package models

import (
	"time"
)

// LoginDataModel 统计用户登录数据 用户学号, 用户姓名,密码,学院,权限,用户token,登录时间
type LoginDataModel struct {
	ID      string    `gorm:"size:20;comment:学号" json:"id"`
	Name    string    `gorm:"size:32;comment:姓名" json:"name"`
	Academy string    `gorm:"size:128;comment:学院" json:"academy"`
	Role    string    `gorm:"size:16;default:2;comment:权限：1管理员,2学生,3老师" json:"role,select(info)"` // 权限  1 管理员  2 学生  3 老师
	Token   string    `gorm:"size:256;comment:token" json:"token"`
	Time    time.Time `gorm:"size:128;comment:time" json:"time"`
}
