package middleware

import (
	"dousheng/cmd/api/conn"
	"strconv"
	"time"
)

func GetUserIdByToken(token string) (int64, error) {
	userIdStr, err := conn.RedisDB.Get(token).Result()
	if err != nil {
		return -1, err
	}
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return -1, err
	}
	return userId, nil
}

func SaveUserId(token string, userId int64) error {
	err := conn.RedisDB.Set(token, userId, time.Duration(6000)*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}
