package service

import (
	"context"
	"dousheng/cmd/favorite/dao"
)

type QueryFavoriteServiceImpl struct {
	ctx context.Context
}

func NewQueryFavoriteServiceImpl(ctx context.Context) *QueryFavoriteServiceImpl {
	return &QueryFavoriteServiceImpl{ctx: ctx}
}

func (receiver *QueryFavoriteServiceImpl) ContainsFavorite(userId int64, videoId int64) (bool, error) {
	favoriteDao := dao.GetFavoriteDaoInstance()
	return favoriteDao.ContainsFavorite(userId, videoId)
}

func (receiver *QueryFavoriteServiceImpl) GetFavoritesByUserId(userId int64) ([]int64, error) {
	favoriteDao := dao.GetFavoriteDaoInstance()
	favorites, err := favoriteDao.GetFavoritesByUserId(userId)
	if err != nil {
		return nil, err
	}
	result := make([]int64, len(favorites))
	for i, f := range favorites {
		result[i] = f.VideoId
	}
	return result, nil
}

func (receiver *QueryFavoriteServiceImpl) CountFavorite(videoId int64) (int64, error) {
	favoriteDao := dao.GetFavoriteDaoInstance()
	return favoriteDao.CountFavorite(videoId)
}
