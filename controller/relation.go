package controller

import (
	"github.com/RaymondCode/simple-demo/conn"
	"github.com/RaymondCode/simple-demo/controller/vo"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserListResponse struct {
	vo.Response
	UserList []string `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	FollowPrefix := "Follow"
	FollowedPrefix := "Followed"
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	actionType := c.Query("action_type")
	claim, err := middleware.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusBadGateway, vo.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid token",
		})
		return
	}

	switch actionType {
	case "1":
		userId := FollowPrefix + claim.UserName
		err = conn.RedisDB.SAdd(userId, toUserId).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, vo.Response{
				StatusCode: 1,
				StatusMsg:  "Unexpected Wrong!",
			})
			return
		}

		toUserId = FollowedPrefix + toUserId
		err = conn.RedisDB.SAdd(toUserId, claim.UserName).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, vo.Response{
				StatusCode: 1,
				StatusMsg:  "Unexpected Wrong!",
			})
			return
		}
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 0,
			StatusMsg:  "关注成功！",
		})

	case "2":
		userId := FollowPrefix + claim.UserName
		err = conn.RedisDB.SRem(userId, toUserId).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, vo.Response{
				StatusCode: 1,
				StatusMsg:  "Unexpected Wrong!",
			})
			return
		}
		toUserId = FollowedPrefix + toUserId
		err = conn.RedisDB.SRem(toUserId, claim.UserName).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, vo.Response{
				StatusCode: 1,
				StatusMsg:  "Unexpected Wrong!",
			})
			return
		}
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 0,
			StatusMsg:  "取关成功！",
		})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	prefix := "Follow"
	userId := c.Query("user_id")
	token := c.Query("token")
	_, err := middleware.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid token",
		})
		return
	}
	userId = prefix + userId
	UserList, err := conn.RedisDB.SMembers(userId).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Response{})
		return
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: vo.Response{StatusCode: 0, StatusMsg: "Load FollowList Success!"},
		UserList: UserList,
	})

	//c.JSON(http.StatusOK, UserListResponse{
	//	Response: vo.Response{
	//		StatusCode: 0,
	//	},
	//	UserList: []vo.User{public.DemoUser},
	//})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	//c.JSON(http.StatusOK, UserListResponse{
	//	Response: vo.Response{
	//		StatusCode: 0,
	//	},
	//	UserList: []vo.User{public.DemoUser},
	//})

	prefix := "Followed"
	userId := c.Query("user_id")
	token := c.Query("token")
	_, err := middleware.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid token",
		})
		return
	}
	userId = prefix + userId
	UserList, err := conn.RedisDB.SMembers(userId).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Response{})
		return
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: vo.Response{StatusCode: 0, StatusMsg: "Load FollowList Success!"},
		UserList: UserList,
	})
}
