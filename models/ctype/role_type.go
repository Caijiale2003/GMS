package ctype

import "encoding/json"

type Role int

const (
	PowerAdmin   Role = 1 // 管理员
	PowerStudent Role = 2 // 学生
	PowerTeacher Role = 3 // 老师
)

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Role) String() string {
	var str string
	switch s {
	case PowerAdmin:
		str = "管理员"
	case PowerStudent:
		str = "学生"
	case PowerTeacher:
		str = "老师"
	default:
		str = "其他"
	}
	return str
}
