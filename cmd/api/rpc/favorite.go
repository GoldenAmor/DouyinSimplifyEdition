package rpc

import (
	"context"
	"dousheng/kitex_gen/favorite"
	"dousheng/kitex_gen/favorite/favoriteservice"
	"dousheng/pkg/conf"
	"errors"
	"github.com/cloudwego/kitex/client"
)

var favoriteClient favoriteservice.Client

func initFavoriteRpc() {
	c, err := favoriteservice.NewClient("example", client.WithHostPorts(conf.FavoriteHostPorts))
	if err != nil {
		panic(err)
	}
	favoriteClient = c
}

func IsFavorite(ctx context.Context, isFavoriteReq *favorite.IsFavoriteReq) (bool, error) {
	resp, err := favoriteClient.IsFavorite(ctx, isFavoriteReq)
	if err != nil {
		return false, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return false, errors.New(resp.BaseResp.StatusMessage)
	}
	return resp.IsFavorite, nil
}

func Like(ctx context.Context, likeReq *favorite.LikeReq) error {
	resp, err := favoriteClient.Like(ctx, likeReq)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errors.New(resp.BaseResp.StatusMessage)
	}
	return nil
}

func UnLike(ctx context.Context, unLikeReq *favorite.UnLikeReq) error {
	resp, err := favoriteClient.UnLike(ctx, unLikeReq)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errors.New(resp.BaseResp.StatusMessage)
	}
	return nil
}

// GetFavoritesByUserId implements the FavoriteServiceImpl interface.
func GetFavoritesByUserId(ctx context.Context, getFavoritesByUserIdReq *favorite.GetFavoritesByUserIdReq) ([]int64, error) {
	resp, err := favoriteClient.GetFavoritesByUserId(ctx, getFavoritesByUserIdReq)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errors.New(resp.BaseResp.StatusMessage)
	}
	return resp.Favorites, nil
}

func CountFavorite(ctx context.Context, countFavoriteReq *favorite.CountFavoriteReq) (int64, error) {
	resp, err := favoriteClient.CountFavorite(ctx, countFavoriteReq)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errors.New(resp.BaseResp.StatusMessage)
	}
	return resp.FavoriteCount, nil
}
