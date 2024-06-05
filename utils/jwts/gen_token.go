package jwts

import (
	"GameManageSystem/global"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

// GenToken 创建 Token
func GenToken(user JwtPayLoad) (string, error) {
	// 将jwt密钥类型转换
	MySecret = []byte(global.Config.Jwt.Secret)
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires))), // 过期时间
			Issuer:    global.Config.Jwt.Issuer,                                                     // 签发人
		},
	}
	// 创建token jwt.NewWithClaims 函数接收两个参数：签名方法（SigningMethodHS256）和声明（claim）
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// 签名：使用指定的密钥对之前创建的声明（claims）进行签名，并将签名后的结果以字符串形式返回
	return token.SignedString(MySecret)
}
