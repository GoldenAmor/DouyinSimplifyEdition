package main

import (
	"context"
	"dousheng/cmd/user/service"
	"dousheng/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// ContainsName implements the UserServiceImpl interface.
func (s *UserServiceImpl) ContainsName(ctx context.Context, containsNameReq *user.ContainsNameReq) (resp *user.ContainsNameResp, err error) {
	// TODO: Your code here...
	username := containsNameReq.Username
	queryUserService := service.NewQueryUserServiceImpl(ctx)
	result, err := queryUserService.ContainsName(username)
	if err != nil {
		resp = &user.ContainsNameResp{
			BaseResp: &user.BaseResp{
				StatusCode:    1,
				StatusMessage: "database error",
			},
		}
		return
	}
	resp = &user.ContainsNameResp{
		BaseResp: &user.BaseResp{
			StatusCode: 0,
		},
		ContainsName: result,
	}
	return
}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, createUserReq *user.CreateUserReq) (resp *user.CreateUserResp, err error) {
	// TODO: Your code here...
	username := createUserReq.Username
	password := createUserReq.Password
	createUserService := service.NewCreateUserServiceImpl(ctx)
	err = createUserService.CreateUser(username, password)
	if err != nil {
		resp = &user.CreateUserResp{
			BaseResp: &user.BaseResp{
				StatusCode:    1,
				StatusMessage: "database error",
			},
		}
		return
	}
	queryUserService := service.NewQueryUserServiceImpl(ctx)
	u, err := queryUserService.GetUserByName(username)
	if err != nil {
		resp = &user.CreateUserResp{
			BaseResp: &user.BaseResp{
				StatusCode:    1,
				StatusMessage: "database error",
			},
		}
		return
	}
	resp = &user.CreateUserResp{
		BaseResp: &user.BaseResp{
			StatusCode: 0,
		},
		UserId: u.ID,
	}
	return
}

// GetUserByName implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserByName(ctx context.Context, getUserByNameReq *user.GetUserByNameReq) (resp *user.GetUserByNameResp, err error) {
	// TODO: Your code here...
	username := getUserByNameReq.Username
	queryUserService := service.NewQueryUserServiceImpl(ctx)
	u, err := queryUserService.GetUserByName(username)
	if err != nil {
		resp = &user.GetUserByNameResp{
			BaseResp: &user.BaseResp{
				StatusCode:    1,
				StatusMessage: "database error",
			},
		}
		return
	}
	resp = &user.GetUserByNameResp{
		BaseResp: &user.BaseResp{
			StatusCode: 0,
		},
		User: &user.User{
			Id:       u.ID,
			Name:     u.Name,
			Password: u.Password,
		},
	}
	return
}

// GetUserById implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserById(ctx context.Context, getUserByIdReq *user.GetUserByIdReq) (resp *user.GetUserByIdResp, err error) {
	// TODO: Your code here...
	id := getUserByIdReq.Id
	queryUserService := service.NewQueryUserServiceImpl(ctx)
	u, err := queryUserService.GetUserById(id)
	if err != nil {
		resp = &user.GetUserByIdResp{
			BaseResp: &user.BaseResp{
				StatusCode:    1,
				StatusMessage: "database error",
			},
		}
		return
	}
	resp = &user.GetUserByIdResp{
		BaseResp: &user.BaseResp{
			StatusCode: 0,
		},
		User: &user.User{
			Id:       u.ID,
			Name:     u.Name,
			Password: u.Password,
		},
	}
	return
}
