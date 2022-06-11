package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/controller/vo"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/repository"
)

func Follow(userId int64, toUserId int64) error {
	lock.Lock()
	defer lock.Unlock()
	if userId == toUserId {
		return errors.New("不能关注自己！")
	}
	toUser := GetUserById(toUserId)
	user := GetUserById(userId)
	FollowDao := dao.GetFollowDaoInstance()
	FollowerDao := dao.GetFollowerDaoInstance()
	record, _, err := FollowDao.GetFollowRecord(userId, toUserId)
	if record.RowsAffected != 0 {
		return errors.New("已关注该用户！")
	}
	if err != nil {
		return err
	}
	err = FollowDao.Create(repository.Follow{
		UserId:   userId,
		ToUserId: toUserId,
	})
	if err != nil {
		return err
	}
	err = FollowerDao.Create(repository.Follower{
		UserId:   toUserId,
		ByUserId: userId,
	})
	if err != nil {
		return err
	}
	//更新关注列表信息与粉丝列表
	err = UpdateUserFollowCount(userId, user.FollowCount+1)
	if err != nil {
		return err
	}
	err = UpdateUserFollowedCount(toUserId, toUser.FollowerCount+1)
	if err != nil {
		return err
	}
	return nil
}

func DeFollow(userId int64, toUserId int64) error {
	lock.Lock()
	defer lock.Unlock()
	toUser := GetUserById(toUserId)
	user := GetUserById(userId)

	//删除关注信息
	FollowDao := dao.GetFollowDaoInstance()
	FollowerDao := dao.GetFollowerDaoInstance()
	record, followRecord, err := FollowDao.GetFollowRecord(userId, toUserId)
	if record.RowsAffected == 0 {
		return errors.New("未关注此用户！")
	}
	if err != nil {
		return err
	}
	err = FollowDao.Delete(repository.Follow{
		Id:       followRecord.Id,
		UserId:   userId,
		ToUserId: toUserId,
	})
	if err != nil {
		return err
	}
	followerRecord, err := FollowerDao.GetFollowerRecord(toUserId, userId)
	if err != nil {
		return err
	}
	err = FollowerDao.Delete(repository.Follower{
		Id:       followerRecord.Id,
		UserId:   toUserId,
		ByUserId: userId,
	})
	if err != nil {
		return err
	}
	//更新关注列表与粉丝列表
	err = UpdateUserFollowCount(userId, user.FollowCount-1)
	if err != nil {
		return err
	}
	err = UpdateUserFollowedCount(toUserId, toUser.FollowerCount-1)
	if err != nil {
		return err
	}
	return nil
}

func GetFollowListByUserID(userId int64) ([]vo.User, error) {
	FollowDao := dao.GetFollowDaoInstance()
	toUsers, err := FollowDao.GetListByUserId(userId)
	if err != nil {
		return nil, err
	}

	result := make([]vo.User, len(toUsers))

	for i, toUser := range toUsers {
		result[i] = *Transform2VoUser(GetUserById(toUser.ToUserId))
	}
	return result, nil
}

func GetFollowerListByUserID(userId int64) ([]vo.User, error) {
	FollowerDao := dao.GetFollowerDaoInstance()
	byUsers, err := FollowerDao.GetListByUserId(userId)
	if err != nil {
		return nil, err
	}
	result := make([]vo.User, len(byUsers))

	for i, toUser := range byUsers {
		result[i] = *Transform2VoUser(GetUserById(toUser.ByUserId))
	}
	return result, nil
}
