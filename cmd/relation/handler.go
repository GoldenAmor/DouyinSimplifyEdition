package main

import (
	"context"
	service "dousheng/cmd/relation/service"
	"dousheng/kitex_gen/relation"
	"fmt"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// CreateRelation implements the UserServiceImpl interface.
func (s *RelationServiceImpl) CreateRelation(ctx context.Context, createRelationReq *relation.CreateRelationReq) (resp *relation.CreateRelationResp, err error) {
	// TODO: Your code here...
	userId := createRelationReq.UserId
	followerId := createRelationReq.FollowerId
	createRelationService := service.NewCreateRelationServiceImpl(ctx)
	err = createRelationService.CreateUser(userId, followerId)
	if err != nil {
		resp = &relation.CreateRelationResp{BaseResp: &relation.BaseResp{
			StatusCode:    1,
			StatusMessage: err.Error(),
		}}
	}
	resp = &relation.CreateRelationResp{BaseResp: &relation.BaseResp{
		StatusCode: 0,
	}}
	return
}

// DeleteRelation implements the UserServiceImpl interface.
func (s *RelationServiceImpl) DeleteRelation(ctx context.Context, deleteRelationReq *relation.DeleteRelationReq) (resp *relation.DeleteRelationResp, err error) {
	// TODO: Your code here...
	userId := deleteRelationReq.UserId
	followerId := deleteRelationReq.FollowerId
	deleteRelationService := service.NewDeleteRelationServiceImpl(ctx)
	err = deleteRelationService.DeleteUser(userId, followerId)
	if err != nil {
		resp = &relation.DeleteRelationResp{BaseResp: &relation.BaseResp{
			StatusCode:    1,
			StatusMessage: err.Error(),
		}}
	}
	resp = &relation.DeleteRelationResp{BaseResp: &relation.BaseResp{
		StatusCode: 0,
	}}
	return
}

// GetFollowers implements the UserServiceImpl interface.
func (s *RelationServiceImpl) GetFollowers(ctx context.Context, getFollowersReq *relation.GetFollowersReq) (resp *relation.GetFollowersResp, err error) {
	// TODO: Your code here...
	userId := getFollowersReq.UserId
	queryRelationService := service.NewQueryRelationServiceImpl(ctx)
	result, err := queryRelationService.GetFollowers(userId)
	if err != nil {
		resp = &relation.GetFollowersResp{BaseResp: &relation.BaseResp{
			StatusCode:    1,
			StatusMessage: err.Error(),
		}}
	}
	resp = &relation.GetFollowersResp{
		BaseResp: &relation.BaseResp{
			StatusCode: 0,
		},
		Followers: result,
	}
	return
}

// GetFollows implements the UserServiceImpl interface.
func (s *RelationServiceImpl) GetFollows(ctx context.Context, getFollowsReq *relation.GetFollowsReq) (resp *relation.GetFollowsResp, err error) {
	// TODO: Your code here...
	userId := getFollowsReq.UserId
	queryRelationService := service.NewQueryRelationServiceImpl(ctx)
	result, err := queryRelationService.GetFollows(userId)
	if err != nil {
		resp = &relation.GetFollowsResp{BaseResp: &relation.BaseResp{
			StatusCode:    1,
			StatusMessage: err.Error(),
		}}
	}
	fmt.Println(userId)
	fmt.Printf("follows:%#v\n", result)
	resp = &relation.GetFollowsResp{
		BaseResp: &relation.BaseResp{
			StatusCode: 0,
		},
		Follows: result,
	}
	return
}

// CountFollowers implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) CountFollowers(ctx context.Context, countFollowersReq *relation.CountFollowersReq) (resp *relation.CountFollowersResp, err error) {
	// TODO: Your code here...
	userId := countFollowersReq.UserId
	queryRelationService := service.NewQueryRelationServiceImpl(ctx)
	followers, err := queryRelationService.GetFollowers(userId)
	if err != nil {
		resp = &relation.CountFollowersResp{BaseResp: &relation.BaseResp{
			StatusCode:    1,
			StatusMessage: err.Error(),
		}}
	}
	resp = &relation.CountFollowersResp{
		BaseResp: &relation.BaseResp{
			StatusCode: 0,
		},
		FollowersCount: int64(len(followers)),
	}
	return
}

// CountFollows implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) CountFollows(ctx context.Context, countFollowsReq *relation.CountFollowsReq) (resp *relation.CountFollowsResp, err error) {
	// TODO: Your code here...
	userId := countFollowsReq.UserId
	queryRelationService := service.NewQueryRelationServiceImpl(ctx)
	follows, err := queryRelationService.GetFollows(userId)
	if err != nil {
		resp = &relation.CountFollowsResp{BaseResp: &relation.BaseResp{
			StatusCode:    1,
			StatusMessage: err.Error(),
		}}
	}
	resp = &relation.CountFollowsResp{
		BaseResp: &relation.BaseResp{
			StatusCode: 0,
		},
		FollowsCount: int64(len(follows)),
	}
	return
}

// IsFollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) IsFollow(ctx context.Context, isFollowReq *relation.IsFollowReq) (resp *relation.IsFollowResp, err error) {
	// TODO: Your code here...
	userId := isFollowReq.UserId
	targetUserId := isFollowReq.TargetUserId
	queryRelationService := service.NewQueryRelationServiceImpl(ctx)
	isFollow, err := queryRelationService.IsFollow(userId, targetUserId)
	if err != nil {
		resp = &relation.IsFollowResp{
			BaseResp: &relation.BaseResp{
				StatusCode: 1,
			},
		}
	}
	resp = &relation.IsFollowResp{
		BaseResp: &relation.BaseResp{
			StatusCode: 0,
		},
		IsFollow: isFollow,
	}
	return
}
