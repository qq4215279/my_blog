/**
 * constants
 * @author liuzhen
 * @Description
 * @version 1.0.0 2021/1/28 16:56
 */

package utils

import (
	"backend/src/global"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var MySecret = []byte(global.Config.Jwt.Secret)

func CreateAccessIdToken(userName string) (string, error) {
	return GenToken(userName, time.Duration(global.Config.Jwt.AccessTokenExpirationTime)*time.Second)
}

func CreateRefreshIdToken(userName string) (string, error) {
	return GenToken(userName, time.Duration(global.Config.Jwt.RefreshTokenExpirationTime)*time.Second)
}

// GenToken 生成JWT
func GenToken(userName string, expirationTime time.Duration) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		userName, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expirationTime).Unix(), // 过期时间
			Issuer:    "my-project",                          // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
