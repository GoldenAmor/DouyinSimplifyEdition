package controller

import (
	"github.com/RaymondCode/simple-demo/controller/vo"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoIdData := c.Query("video_id")
	actionType := c.Query("action_type")
	videoId, err := strconv.ParseInt(videoIdData, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "video_id error"})
		return
	}
	//鉴权
	user, err := service.GetUserByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	//点赞
	if actionType == "1" {
		err = service.Like(user.ID, videoId)
	}
	//取消点赞
	if actionType == "2" {
		err = service.UnLike(user.ID, videoId)
	}
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "database error"})
		return
	}
	c.JSON(http.StatusOK, vo.Response{StatusCode: 0})
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
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
	//鉴权
	_, err = service.GetUserByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	//获取列表
	result, err := service.GetFavoritesByUserId(userId)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "database error"})
		return
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: vo.Response{
			StatusCode: 0,
		},
		VideoList: result,
	})
}
