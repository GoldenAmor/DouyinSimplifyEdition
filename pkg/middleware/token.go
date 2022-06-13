package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserName string
	PassWord string
	jwt.StandardClaims
}

func GenerateToken(nameArg string, passwordArg string) string {
	nowTime := time.Now()
	expireTime := nowTime.Add(300 * time.Second)
	issuer := "wangfeng"
	claims := Claims{
		UserName: nameArg,
		PassWord: passwordArg,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("golang"))
	return token
}
