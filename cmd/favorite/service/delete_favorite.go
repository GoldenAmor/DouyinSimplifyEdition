package service

import (
	"context"
	"dousheng/cmd/favorite/dao"
)

type DeleteFavoriteServiceImpl struct {
	ctx context.Context
}

func NewDeleteFavoriteServiceImpl(ctx context.Context) *DeleteFavoriteServiceImpl {
	return &DeleteFavoriteServiceImpl{ctx: ctx}
}

func (receiver *DeleteFavoriteServiceImpl) DeleteFavorite(userId int64, videoId int64) error {
	favoriteDao := dao.GetFavoriteDaoInstance()
	v, err := favoriteDao.GetFavoriteByUserIdAndVideoId(userId, videoId)
	if err != nil {
		return err
	}
	return favoriteDao.DeleteFavorite(v)
}
