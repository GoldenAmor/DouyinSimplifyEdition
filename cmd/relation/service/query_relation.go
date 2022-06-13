package service

import (
	"context"
	"dousheng/cmd/relation/dao"
)

type QueryRelationServiceImpl struct {
	ctx context.Context
}

func NewQueryRelationServiceImpl(ctx context.Context) *QueryRelationServiceImpl {
	return &QueryRelationServiceImpl{ctx: ctx}
}

func (receiver *QueryRelationServiceImpl) GetFollows(userId int64) ([]int64, error) {
	relationDao := dao.GetRelationDaoInstance()
	relations, err := relationDao.GetFollows(userId)
	if err != nil {
		return nil, err
	}
	result := make([]int64, len(relations))
	for i, relation := range relations {
		result[i] = relation.UserId
	}
	return result, nil
}

func (receiver *QueryRelationServiceImpl) GetFollowers(userId int64) ([]int64, error) {
	relationDao := dao.GetRelationDaoInstance()
	relations, err := relationDao.GetFollowers(userId)
	if err != nil {
		return nil, err
	}
	result := make([]int64, len(relations))
	for i, relation := range relations {
		result[i] = relation.FollowerId
	}
	return result, nil
}
func (receiver *QueryRelationServiceImpl) IsFollow(userId int64, targetUserId int64) (bool, error) {
	relationDao := dao.GetRelationDaoInstance()
	isFollow, err := relationDao.IsFollow(targetUserId, userId)
	if err != nil {
		return false, err
	}
	return isFollow, nil
}
