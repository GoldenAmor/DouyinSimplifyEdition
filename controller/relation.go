package controller

import (
	"github.com/RaymondCode/simple-demo/controller/vo"
	"github.com/RaymondCode/simple-demo/service"
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
	actionType := c.Query("action_type")
	toUserId, err := strconv.ParseInt(toUserIdData, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "video_id error"})
		return
	}
	user, err := service.GetUserByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	if actionType == "1" {
		err = service.Follow(user.ID, toUserId)
	}

	if actionType == "2" {
		err = service.DeFollow(user.ID, toUserId)
	}
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "database error"})
		return
	}
	c.JSON(http.StatusOK, vo.Response{StatusCode: 0})
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {

	token := c.Query("token")
	userIdData := c.Query("user_id")
	_, err := service.GetUserByToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Response{StatusCode: 1, StatusMsg: "Invalid token!"})
		return
	}
	userId, err := strconv.ParseInt(userIdData, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Response{StatusCode: 1, StatusMsg: "Unexpected error!"})
		return
	}
	result, err := service.GetFollowListByUserID(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Response{StatusCode: 1, StatusMsg: "Unexpected error!"})
		return
	}

	c.JSON(http.StatusOK, UserListResponse{
		Response: vo.Response{
			StatusCode: 0,
			StatusMsg:  "获取关注列表成功！",
		},

		UserList: result,
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {

	token := c.Query("token")
	userIdData := c.Query("user_id")
	_, err := service.GetUserByToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Response{StatusCode: 1, StatusMsg: "Invalid token!"})
		return
	}
	userId, err := strconv.ParseInt(userIdData, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Response{StatusCode: 1, StatusMsg: "Unexpected error!"})
		return
	}
	result, err := service.GetFollowerListByUserID(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Response{StatusCode: 1, StatusMsg: "Unexpected error!"})
		return
	}

	c.JSON(http.StatusOK, UserListResponse{
		Response: vo.Response{
			StatusCode: 0,
			StatusMsg:  "获取粉丝列表成功！",
		},
		UserList: result,
	})
}
