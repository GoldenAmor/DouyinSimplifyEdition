package controller

import (
	"context"
	"dousheng/cmd/api/controller/vo"
	"dousheng/cmd/api/rpc"
	"dousheng/cmd/api/utils"
	"dousheng/kitex_gen/relation"
	"dousheng/kitex_gen/user"
	"dousheng/pkg/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	vo.Response
	UserList []vo.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserIdData := c.Query("to_user_id")
	toUserId, err := strconv.ParseInt(toUserIdData, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "to_user_id error"})
		return
	}
	actionType := c.Query("action_type")
	uId, err := middleware.GetUserIdByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "Invalid Token"})
		return
	}
	if actionType == "1" {
		err = rpc.CreateRelation(context.Background(), &relation.CreateRelationReq{
			UserId:     toUserId,
			FollowerId: uId,
		})
	}
	if actionType == "2" {
		err = rpc.DeleteRelation(context.Background(), &relation.DeleteRelationReq{
			UserId:     toUserId,
			FollowerId: uId,
		})
	}
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: err.Error()})
	}
	c.JSON(http.StatusOK, vo.Response{StatusCode: 0})
}

// FollowList all users have same list
func FollowList(c *gin.Context) {
	userIdData := c.Query("user_id")
	userId, err := strconv.ParseInt(userIdData, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid user_id",
		})
		return
	}
	token := c.Query("token")
	uId, err := middleware.GetUserIdByToken(token)
	if err != nil {
		uId = -1
	}
	follows, err := rpc.GetFollows(context.Background(), &relation.GetFollowsReq{UserId: userId})
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	fmt.Printf("follows:%#v\n", follows)
	userList := make([]vo.User, len(follows))
	for i, f := range follows {
		u, err := rpc.GetUserById(context.Background(), &user.GetUserByIdReq{Id: f})
		if err != nil {
			c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "Invalid Token"})
			return
		}
		us, err := utils.PackageUser(uId, u)
		if err != nil {
			c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "Invalid Token"})
			return
		}
		userList[i] = *us
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: vo.Response{StatusCode: 0, StatusMsg: "Load FollowList Success!"},
		UserList: userList,
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	userIdData := c.Query("user_id")
	userId, err := strconv.ParseInt(userIdData, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid user_id",
		})
		return
	}
	token := c.Query("token")
	uId, err := middleware.GetUserIdByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "Invalid Token"})
		return
	}
	followers, err := rpc.GetFollowers(context.Background(), &relation.GetFollowersReq{UserId: userId})
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	userList := make([]vo.User, len(followers))
	for i, f := range followers {
		u, err := rpc.GetUserById(context.Background(), &user.GetUserByIdReq{Id: f})
		if err != nil {
			c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		us, err := utils.PackageUser(uId, u)
		if err != nil {
			c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "Invalid Token"})
			return
		}
		userList[i] = *us
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: vo.Response{StatusCode: 0, StatusMsg: "Load FollowerList Success!"},
		UserList: userList,
	})
}
