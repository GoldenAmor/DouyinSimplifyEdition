package main

import (
	"context"
	"dousheng/cmd/favorite/service"
	"dousheng/kitex_gen/favorite"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// IsFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) IsFavorite(ctx context.Context, isFavoriteReq *favorite.IsFavoriteReq) (resp *favorite.IsFavoriteResp, err error) {
	// TODO: Your code here...
	userId := isFavoriteReq.UserId
	videoId := isFavoriteReq.VideoId
	queryFavoriteService := service.NewQueryFavoriteServiceImpl(ctx)
	containsFavorite, err := queryFavoriteService.ContainsFavorite(userId, videoId)
	if err != nil {
		resp = &favorite.IsFavoriteResp{
			BaseResp: &favorite.BaseResp{
				StatusCode:    1,
				StatusMessage: err.Error(),
			},
		}
		return
	}
	resp = &favorite.IsFavoriteResp{
		BaseResp:   &favorite.BaseResp{StatusCode: 0},
		IsFavorite: containsFavorite,
	}
	return
}

// Like implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) Like(ctx context.Context, likeReq *favorite.LikeReq) (resp *favorite.LikeResp, err error) {
	// TODO: Your code here...
	userId := likeReq.UserId
	videoId := likeReq.VideoId
	createFavoriteService := service.NewCreateFavoriteServiceImpl(ctx)
	err = createFavoriteService.CreateFavorite(userId, videoId)
	if err != nil {
		resp = &favorite.LikeResp{BaseResp: &favorite.BaseResp{
			StatusCode:    1,
			StatusMessage: err.Error(),
		}}
		return
	}
	resp = &favorite.LikeResp{BaseResp: &favorite.BaseResp{
		StatusCode: 0,
	}}
	return
}

// UnLike implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) UnLike(ctx context.Context, unLikeReq *favorite.UnLikeReq) (resp *favorite.UnLikeResp, err error) {
	// TODO: Your code here...
	userId := unLikeReq.UserId
	videoId := unLikeReq.VideoId
	deleteFavoriteService := service.NewDeleteFavoriteServiceImpl(ctx)
	err = deleteFavoriteService.DeleteFavorite(userId, videoId)
	if err != nil {
		resp = &favorite.UnLikeResp{BaseResp: &favorite.BaseResp{
			StatusCode:    1,
			StatusMessage: err.Error(),
		}}
		return
	}
	resp = &favorite.UnLikeResp{BaseResp: &favorite.BaseResp{
		StatusCode: 0,
	}}
	return
}

// GetFavoritesByUserId implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) GetFavoritesByUserId(ctx context.Context, getFavoritesByUserIdReq *favorite.GetFavoritesByUserIdReq) (resp *favorite.GetFavoritesByUserIdResp, err error) {
	// TODO: Your code here...
	userId := getFavoritesByUserIdReq.UserId
	queryFavoriteService := service.NewQueryFavoriteServiceImpl(ctx)
	videoIdList, err := queryFavoriteService.GetFavoritesByUserId(userId)
	if err != nil {
		resp = &favorite.GetFavoritesByUserIdResp{
			BaseResp: &favorite.BaseResp{
				StatusCode:    1,
				StatusMessage: err.Error(),
			},
		}
		return
	}
	resp = &favorite.GetFavoritesByUserIdResp{
		BaseResp: &favorite.BaseResp{
			StatusCode: 0,
		},
		Favorites: videoIdList,
	}
	return
}

// CountFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) CountFavorite(ctx context.Context, countFavoriteReq *favorite.CountFavoriteReq) (resp *favorite.CountFavoriteResp, err error) {
	// TODO: Your code here...
	videoId := countFavoriteReq.VideoId
	queryFavoriteService := service.NewQueryFavoriteServiceImpl(ctx)
	favoriteCount, err := queryFavoriteService.CountFavorite(videoId)
	if err != nil {
		resp = &favorite.CountFavoriteResp{
			BaseResp: &favorite.BaseResp{
				StatusCode:    1,
				StatusMessage: err.Error(),
			},
		}
		return
	}
	resp = &favorite.CountFavoriteResp{
		BaseResp:      &favorite.BaseResp{StatusCode: 0},
		FavoriteCount: favoriteCount,
	}
	return
}
