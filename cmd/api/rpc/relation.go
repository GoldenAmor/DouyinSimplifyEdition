package rpc

import (
	"context"
	"dousheng/kitex_gen/relation"
	"dousheng/kitex_gen/relation/relationservice"
	"dousheng/pkg/conf"
	"errors"
	"github.com/cloudwego/kitex/client"
)

var relationClient relationservice.Client

func initRelationRpc() {
	c, err := relationservice.NewClient("example", client.WithHostPorts(conf.RelationHostPorts))
	if err != nil {
		panic(err)
	}
	relationClient = c
}

func CreateRelation(ctx context.Context, createRelationReq *relation.CreateRelationReq) error {
	resp, err := relationClient.CreateRelation(ctx, createRelationReq)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errors.New(resp.BaseResp.StatusMessage)
	}
	return nil
}

func DeleteRelation(ctx context.Context, deleteRelationReq *relation.DeleteRelationReq) error {
	resp, err := relationClient.DeleteRelation(ctx, deleteRelationReq)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errors.New(resp.BaseResp.StatusMessage)
	}
	return nil
}

func GetFollowers(ctx context.Context, getFollowersReq *relation.GetFollowersReq) ([]int64, error) {
	resp, err := relationClient.GetFollowers(ctx, getFollowersReq)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errors.New(resp.BaseResp.StatusMessage)
	}
	return resp.Followers, nil
}

func GetFollows(ctx context.Context, getFollowsReq *relation.GetFollowsReq) ([]int64, error) {
	resp, err := relationClient.GetFollows(ctx, getFollowsReq)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errors.New(resp.BaseResp.StatusMessage)
	}
	return resp.Follows, nil
}

func CountFollowers(ctx context.Context, countFollowersReq *relation.CountFollowersReq) (int64, error) {
	resp, err := relationClient.CountFollowers(ctx, countFollowersReq)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errors.New(resp.BaseResp.StatusMessage)
	}
	return resp.FollowersCount, nil
}

func CountFollows(ctx context.Context, countFollowsReq *relation.CountFollowsReq) (int64, error) {
	resp, err := relationClient.CountFollows(ctx, countFollowsReq)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errors.New(resp.BaseResp.StatusMessage)
	}
	return resp.FollowsCount, nil
}
func IsFollow(ctx context.Context, isFollowReq *relation.IsFollowReq) (bool, error) {
	resp, err := relationClient.IsFollow(ctx, isFollowReq)
	if err != nil {
		return false, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return false, errors.New(resp.BaseResp.StatusMessage)
	}
	return resp.IsFollow, nil
}
