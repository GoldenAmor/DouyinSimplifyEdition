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

func GenerateToken(nameArg string, passwordArg string) (string, error) {
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

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("golang"))
	return token, err
}
func ParseToken(token string) error {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("golang"), nil
	})
	if err != nil {
		return err
	}
	if tokenClaims != nil {
		if _, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return nil
		}
	}
	return err
}
