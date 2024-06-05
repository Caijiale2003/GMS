package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
)

// JwtPayLoad jwt中payload数据
type JwtPayLoad struct {
	UserID   string `json:"user_id"`  // 用户id
	Username string `json:"username"` // 用户名
	Role     int    `json:"role"`     // 权限  1 管理员  2 学生  3 老师

}

var MySecret []byte

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}
