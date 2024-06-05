package user_api

import (
	"GameManageSystem/global"
	"GameManageSystem/models"
	"GameManageSystem/models/ctype"
	"GameManageSystem/models/res"
	"GameManageSystem/utils/pwd"
	"github.com/gin-gonic/gin"
)

type UserLoginRequest struct {
	ID       string     `json:"id" binding:"required" msg:"请输入学号"`
	Name     string     `json:"name" binding:"required" msg:"请输入姓名"`
	Password string     `json:"password" binding:"required" msg:"请输入密码"`
	Gender   string     `json:"gender" binding:"required" msg:"请输入性别"`
	Academy  string     `json:"academy" binding:"required" msg:"请输入学院"`
	Major    string     `json:"major" binding:"required" msg:"请输入专业"`
	Role     ctype.Role `json:"role" binding:"required" msg:"请输入专业"`
}

func (UserApi) UserLoginView(c *gin.Context) {
	// 参数绑定校验
	var re UserLoginRequest
	err := c.ShouldBindJSON(&re)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 校验数据库中是否存在该用户
	var user models.UserModel
	err = global.DB.Take(&user, "id = ?", re.ID).Error
	if err == nil {
		res.FailWithMessage("该用户已存在，请直接登录！", c)
		return
	}

	// 判断身份:既不是学生也不是老师直接限定为学生
	if re.Role != ctype.PowerStudent && re.Role != ctype.PowerAdmin {
		re.Role = ctype.PowerStudent
	}
	// 写入数据库
	hashPwd := pwd.HashPwd(re.Password)
	user = models.UserModel{
		ID:       re.ID,
		Name:     re.Name,
		Password: hashPwd,
		Academy:  re.Academy,
		Role:     re.Role,
		Gender:   re.Gender,
		Major:    re.Major,
	}
	err = global.DB.Create(&user).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("注册失败", c)
		return
	}
	res.OkWithData("注册成功", c)
}
