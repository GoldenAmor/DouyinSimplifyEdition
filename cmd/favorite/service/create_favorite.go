package service

import (
	"context"
	"dousheng/cmd/favorite/dao"
	"dousheng/cmd/favorite/repository"
)

type CreateFavoriteServiceImpl struct {
	ctx context.Context
}

func NewCreateFavoriteServiceImpl(ctx context.Context) *CreateFavoriteServiceImpl {
	return &CreateFavoriteServiceImpl{ctx: ctx}
}

func (receiver *CreateFavoriteServiceImpl) CreateFavorite(userId int64, videoId int64) error {
	favoriteDao := dao.GetFavoriteDaoInstance()
	return favoriteDao.CreateFavorite(&repository.Favorite{
		UserId:  userId,
		VideoId: videoId,
	})
}
