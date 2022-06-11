package dao

import (
	"github.com/RaymondCode/simple-demo/conn"
	"github.com/RaymondCode/simple-demo/repository"
	"sync"
)

type User struct{}

var (
	user     User
	userOnce sync.Once
)

func GetUserDaoInstance() *User {
	userOnce.Do(func() {
		user = User{}
	})
	return &user
}

func (u *User) Create(object *repository.User) error {
	return conn.DB.Create(object).Error
}

func (u *User) GetByName(username string) *repository.User {
	result := &repository.User{}
	err := conn.DB.Where("name = ?", username).First(result).Error
	if err != nil {
		return nil
	}
	return result
}

func (u *User) GetById(id int64) *repository.User {
	result := &repository.User{}
	err := conn.DB.Where("id = ?", id).First(result).Error
	if err != nil {
		return nil
	}
	return result
}

func (u *User) UpdateUserFollowCount(toUserId int64, count int64) error {
	return conn.DB.Model(&User{}).Where("id = ?", toUserId).Update("follow_count", count).Error
}

func (u *User) UpdateUserFollowedCount(byUserId int64, count int64) error {
	return conn.DB.Model(&User{}).Where("id = ?", byUserId).Update("follower_count", count).Error
}
