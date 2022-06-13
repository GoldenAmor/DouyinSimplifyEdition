package service

import (
	"context"
	"dousheng/cmd/relation/dao"
	"dousheng/cmd/relation/repository"
)

type DeleteRelationServiceImpl struct {
	ctx context.Context
}

func NewDeleteRelationServiceImpl(ctx context.Context) *DeleteRelationServiceImpl {
	return &DeleteRelationServiceImpl{ctx: ctx}
}

func (receiver *DeleteRelationServiceImpl) DeleteUser(userId int64, followerId int64) error {
	relationDao := dao.GetRelationDaoInstance()
	result, err := relationDao.GetRelation(userId, followerId)
	if err != nil {
		return err
	}
	return relationDao.DeleteRelation(repository.Relation{
		Id: result.Id,
	})
}
