package service

import (
	"context"
	"dousheng/cmd/relation/dao"
	"dousheng/cmd/relation/repository"
)

type CreateRelationServiceImpl struct {
	ctx context.Context
}

func NewCreateRelationServiceImpl(ctx context.Context) *CreateRelationServiceImpl {
	return &CreateRelationServiceImpl{ctx: ctx}
}

func (receiver *CreateRelationServiceImpl) CreateUser(userId int64, followerId int64) error {
	relationDao := dao.GetRelationDaoInstance()
	return relationDao.CreateRelation(repository.Relation{
		UserId:     userId,
		FollowerId: followerId,
	})
}
