package dao

import (
	"github.com/RaymondCode/simple-demo/conn"
	"github.com/RaymondCode/simple-demo/repository"
	"gorm.io/gorm"
	"sync"
)

type Follow struct{}

var (
	follow     Follow
	followOnce sync.Once
)

func GetFollowDaoInstance() *Follow {
	followOnce.Do(func() {
		follow = Follow{}
	})
	return &follow
}

func (c *Follow) Create(object repository.Follow) error {
	return conn.DB.Create(&object).Error
}

func (c *Follow) Delete(object repository.Follow) error {
	return conn.DB.Delete(&object).Error
}

func (c *Follow) GetListByUserId(userId int64) ([]repository.Follow, error) {
	var result []repository.Follow
	err := conn.DB.Order("created_at desc").Where("user_id = ?", userId).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Follow) GetFollowRecord(userId int64, toUserId int64) (*gorm.DB, repository.Follow, error) {
	var result repository.Follow
	record := conn.DB.Order("created_at desc").Where("user_id = ? AND to_user_id = ?", userId, toUserId).Find(&result)
	if record.Error != nil {
		return record, repository.Follow{}, record.Error
	}
	return record, result, nil
}
