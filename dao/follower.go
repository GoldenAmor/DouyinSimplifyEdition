package dao

import (
	"github.com/RaymondCode/simple-demo/conn"
	"github.com/RaymondCode/simple-demo/repository"
	"sync"
)

type Follower struct{}

var (
	follower     Follower
	followerOnce sync.Once
)

func GetFollowerDaoInstance() *Follower {
	followerOnce.Do(func() {
		follower = Follower{}
	})
	return &follower
}

func (c *Follower) Create(object repository.Follower) error {
	return conn.DB.Create(&object).Error
}

func (c *Follower) Delete(object repository.Follower) error {
	return conn.DB.Delete(&object).Error
}

func (c *Follower) GetListByUserId(userId int64) ([]repository.Follower, error) {
	var result []repository.Follower
	err := conn.DB.Order("created_at desc").Where("user_id = ?", userId).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Follower) GetFollowerRecord(userId int64, byUserId int64) (repository.Follower, error) {
	var result repository.Follower
	err := conn.DB.Order("created_at desc").Where("user_id = ? AND by_user_id = ?", userId, byUserId).Find(&result).Error
	if err != nil {
		return repository.Follower{}, err
	}
	return result, nil
}
