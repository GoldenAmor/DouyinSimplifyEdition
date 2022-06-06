package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/controller/vo"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync/atomic"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin

var userIdSequence = int64(1)

type UserLoginResponse struct {
	vo.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	vo.Response
	User vo.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	//检查用户名是否重复
	containsName := service.ContainsName(username)
	if containsName {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  "User already existed",
		})
		return
	}
	//自增主键
	atomic.AddInt64(&userIdSequence, 1)
	//插入记录
	createError := service.CreateUser(username, password)
	if createError != nil {
		fmt.Println(createError.Error())
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  "Unexpected fault!",
		})
		return
	}
	//合成token
	token, err := middleware.GenerateToken(username, password)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  "Fail to generate token",
		})
		return
	}
	//获取id
	id := service.GetUserByName(username).ID
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: vo.Response{StatusMsg: "Register success!"},
		UserId:   id,
		Token:    token,
	})
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//检查用户名是否存在
	u := service.GetUserByName(username)
	if u == nil {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist",
		})
		return
	}
	//检查密码是否正确
	if password != u.Password {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  "Wrong password!",
		})
		return
	}
	token, err := middleware.GenerateToken(username, password)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  "Fail to generate token",
		})
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: vo.Response{StatusMsg: "Login Success!"},
		UserId:   u.ID,
		Token:    token,
	})
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	userIdData := c.Query("user_id")
	userId, err := strconv.ParseInt(userIdData, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid user_id",
		})
		return
	}
	//鉴权并获取登录用户信息
	_, err = middleware.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid token",
		})
		return
	}
	user, err := service.GetUserByToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid token",
		})
		return
	}
	//获取目标用户信息
	userTarget := service.GetUserById(userId)
	//当前用户是否关注了该用户
	isFollow := service.IsFollow(user.ID, userTarget.ID)
	//转换
	u := service.Transform2VoUser(user)
	u.IsFollow = isFollow
	c.JSON(http.StatusOK, UserResponse{
		Response: vo.Response{StatusCode: 0},
		User:     *u,
	})

}
