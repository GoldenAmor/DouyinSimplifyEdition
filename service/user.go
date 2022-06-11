package service

import (
	"fmt"
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

func UpdateUserFollowCount(toUserId int64, count int64) error {
	fmt.Println("更新")
	fmt.Println(count)
	userDao := dao.GetUserDaoInstance()
	return userDao.UpdateUserFollowCount(toUserId, count)
}

func UpdateUserFollowedCount(byUserId int64, count int64) error {
	fmt.Println("更新")
	fmt.Println(count)
	userDao := dao.GetUserDaoInstance()
	return userDao.UpdateUserFollowedCount(byUserId, count)
}

func IsFollow(userId int64, targetId int64) (bool, error) {
	followDao := dao.GetFollowDaoInstance()
	record, _, err := followDao.GetFollowRecord(userId, targetId)
	if err != nil {
		return false, err
	}
	if record.RowsAffected == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func Transform2VoUser(user *repository.User) *vo.User {
	return &vo.User{
		Id:            user.ID,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
	}
}
