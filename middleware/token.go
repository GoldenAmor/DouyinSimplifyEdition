package middleware

import (
	"encoding/json"
	"github.com/RaymondCode/simple-demo/service"
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
	claims := Claims{
		UserName: nameArg,
		PassWord: passwordArg,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("golang"))
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("golang"), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func CacheToken(token string) error {
	claims, err := ParseToken(token)
	if err != nil {
		return err
	}
	user := service.GetUserByName(claims.UserName)
	userInfoTemp, err := json.Marshal(&user)
	if err != nil {
		return err
	}
	userInfo := string(userInfoTemp)
	timeForLive := time.Duration(claims.ExpiresAt) * time.Second
	err = service.SaveUser2Redis(token, userInfo, timeForLive)
	if err != nil {
		return err
	}
	return nil
}
