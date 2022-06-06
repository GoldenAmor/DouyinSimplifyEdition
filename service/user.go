package service

import (
	"github.com/RaymondCode/simple-demo/controller/vo"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/repository"
)

func ContainsName(username string) bool {
	userDao := dao.GetUserDaoInstance()
	user := userDao.GetByName(username)
	return nil != user
}

func CreateUser(username string, password string) error {
	userDao := dao.GetUserDaoInstance()
	return userDao.Create(&repository.User{
		Name:          username,
		Password:      password,
		FollowCount:   0,
		FollowerCount: 0,
	})
}

func GetUserByName(username string) *repository.User {
	userDao := dao.GetUserDaoInstance()
	return userDao.GetByName(username)
}

func GetUserById(id int64) *repository.User {
	userDao := dao.GetUserDaoInstance()
	return userDao.GetById(id)
}

func IsFollow(userId int64, targetId int64) bool {
	return true
}

func Transform2VoUser(user *repository.User) *vo.User {
	return &vo.User{
		Id:            user.ID,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
	}
}
