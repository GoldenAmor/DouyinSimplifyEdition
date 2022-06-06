package service

import (
	"github.com/RaymondCode/simple-demo/repository"
)

func GetUserByToken(token string) (*repository.User, error) {
	return &repository.User{
		ID: 1,
	}, nil
}
