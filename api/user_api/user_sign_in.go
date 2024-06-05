package user_api

import (
	"GameManageSystem/global"
	"GameManageSystem/models"
	"GameManageSystem/models/res"
	"GameManageSystem/utils/jwts"
	"GameManageSystem/utils/pwd"
	"github.com/gin-gonic/gin"
	"time"
)

type UserSignRequest struct {
	ID       string `json:"id" binding:"required" msg:"请输入学号"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

func (UserApi) UserSignView(c *gin.Context) {
	// 参数绑定校验
	var re UserSignRequest
	err := c.ShouldBindJSON(&re)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 校验数据库中是否存在该用户
	var user models.UserModel
	err = global.DB.Take(&user, "id = ?", re.ID).Error
	if err != nil {
		res.FailWithMessage("该用户不存在，请先注册！", c)
		return
	}

	// 校验密码
	isCheck := pwd.CheckPwd(user.Password, re.Password)
	if !isCheck {
		// 后台日志警告
		global.Log.Warn("用户名密码错误")
		// 返回错误
		res.FailWithMessage("用户名或密码错误", c)
		return
	}

	// 登录操作
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		Username: user.Name,
		Role:     int(user.Role),
		UserID:   user.ID,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("token生成失败", c)
		return
	}
	c.Request.Header.Set("token", token)

	global.DB.Create(&models.LoginDataModel{
		ID:      user.ID,
		Name:    user.Name,
		Academy: user.Academy,
		Role:    user.Role.String(),
		Token:   token,
		Time:    time.Now(),
	})
	res.OkWithData(token, c)
}
