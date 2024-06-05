package jwts

import (
	"GameManageSystem/global"
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
)

// ParseToken 解析 token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	// 将jwt密钥类型转换
	MySecret = []byte(global.Config.Jwt.Secret)
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无效 token")
}
