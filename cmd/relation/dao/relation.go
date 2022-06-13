package dao

import (
	"dousheng/cmd/relation/conn"
	"dousheng/cmd/relation/repository"
	"sync"
)

type RelationDao struct{}

var (
	relationDao  RelationDao
	relationOnce sync.Once
)

func GetRelationDaoInstance() *RelationDao {
	relationOnce.Do(func() {
		relationDao = RelationDao{}
	})
	return &relationDao
}

func (receiver *RelationDao) GetRelation(userId int64, followerId int64) (*repository.Relation, error) {
	result := &repository.Relation{}
	err := conn.DB.Where("user_id = ? and follower_id =?", userId, followerId).First(result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (receiver *RelationDao) GetFollows(userId int64) ([]repository.Relation, error) {
	var result []repository.Relation
	err := conn.DB.Where("follower_id =?", userId).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (receiver *RelationDao) GetFollowers(userId int64) ([]repository.Relation, error) {
	var result []repository.Relation
	err := conn.DB.Where("user_id =?", userId).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (receiver *RelationDao) CreateRelation(object repository.Relation) error {
	return conn.DB.Create(&object).Error
}

func (receiver *RelationDao) DeleteRelation(object repository.Relation) error {
	return conn.DB.Delete(&object).Error
}

func (receiver *RelationDao) GetFollowerCount(userId int64) (int64, error) {
	var count int64
	err := conn.DB.Model(&repository.Relation{}).Where("user_id = ?", userId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (receiver *RelationDao) IsFollow(userId int64, followerId int64) (bool, error) {
	var count int64
	err := conn.DB.Model(&repository.Relation{}).Where("user_id = ? and follower_id = ?", userId, followerId).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count != 0, nil
}

func (receiver *RelationDao) GetFollowCount(userId int64) (int64, error) {
	var count int64
	err := conn.DB.Model(&repository.Relation{}).Select("count(1) as count").Where("follower_id = ?", userId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
