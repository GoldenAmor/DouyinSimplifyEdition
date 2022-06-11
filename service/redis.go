package service

import (
	"encoding/json"
	"github.com/RaymondCode/simple-demo/conn"
	"github.com/RaymondCode/simple-demo/repository"
	"time"
)

func GetUserByToken(token string) (*repository.User, error) {
	users, err := conn.RedisDB.Get(token).Result()
	if err != nil {
		return nil, err
	}
	user := repository.User{}
	err = json.Unmarshal([]byte(users), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func SaveUser2Redis(token string, userInfo string, expireTime time.Duration) error {
	err := conn.RedisDB.Set(token, userInfo, expireTime).Err()
	if err != nil {
		return err
	}
	return nil
}
