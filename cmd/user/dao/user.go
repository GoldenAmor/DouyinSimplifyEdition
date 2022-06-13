package dao

import (
	"dousheng/cmd/user/conn"
	"dousheng/cmd/user/repository"
	"sync"
)

type UserDao struct{}

var (
	userDao  UserDao
	userOnce sync.Once
)

func GetUserDaoInstance() *UserDao {
	userOnce.Do(func() {
		userDao = UserDao{}
	})
	return &userDao
}

func (u *UserDao) CreateUser(object *repository.User) error {
	return conn.DB.Create(object).Error
}

func (u *UserDao) GetUserByName(username string) (*repository.User, error) {
	result := &repository.User{}
	err := conn.DB.Where("name = ?", username).First(result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserDao) GetUserById(id int64) (*repository.User, error) {
	result := &repository.User{}
	err := conn.DB.Where("id = ?", id).First(result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserDao) ContainsName(username string) (bool, error) {
	var count int64
	err := conn.DB.Model(&repository.User{}).Where("name = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}
