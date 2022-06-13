package service

import (
	"context"
	"dousheng/cmd/user/dao"
	"dousheng/cmd/user/repository"
)

type QueryUserServiceImpl struct {
	ctx context.Context
}

func NewQueryUserServiceImpl(ctx context.Context) *QueryUserServiceImpl {
	return &QueryUserServiceImpl{ctx: ctx}
}

func (receiver *QueryUserServiceImpl) ContainsName(username string) (bool, error) {
	userDao := dao.GetUserDaoInstance()
	return userDao.ContainsName(username)
}

func (receiver *QueryUserServiceImpl) GetUserByName(username string) (*repository.User, error) {
	userDao := dao.GetUserDaoInstance()
	return userDao.GetUserByName(username)
}

func (receiver *QueryUserServiceImpl) GetUserById(id int64) (*repository.User, error) {
	userDao := dao.GetUserDaoInstance()
	return userDao.GetUserById(id)
}
