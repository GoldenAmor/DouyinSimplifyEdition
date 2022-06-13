package rpc

import (
	"context"
	"dousheng/kitex_gen/user"
	"dousheng/kitex_gen/user/userservice"
	"dousheng/pkg/conf"
	"errors"
	"github.com/cloudwego/kitex/client"
)

var userClient userservice.Client

func initUserRpc() {
	c, err := userservice.NewClient("example", client.WithHostPorts(conf.UserHostPorts))
	if err != nil {
		panic(err)
	}
	userClient = c
}

func ContainsName(ctx context.Context, containsNameReq *user.ContainsNameReq) (bool, error) {
	resp, err := userClient.ContainsName(ctx, containsNameReq)
	if err != nil {
		return false, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return false, errors.New(resp.BaseResp.StatusMessage)
	}
	return resp.ContainsName, nil
}

func CreateUser(ctx context.Context, createUserReq *user.CreateUserReq) (int64, error) {
	resp, err := userClient.CreateUser(ctx, createUserReq)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errors.New(resp.BaseResp.StatusMessage)
	}
	return resp.UserId, nil
}

func GetUserByName(ctx context.Context, getUserByNameReq *user.GetUserByNameReq) (*user.User, error) {
	resp, err := userClient.GetUserByName(ctx, getUserByNameReq)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errors.New(resp.BaseResp.StatusMessage)
	}
	return resp.User, nil
}

func GetUserById(ctx context.Context, getUserByIdReq *user.GetUserByIdReq) (*user.User, error) {
	resp, err := userClient.GetUserById(ctx, getUserByIdReq)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errors.New(resp.BaseResp.StatusMessage)
	}
	return resp.User, nil
}
