package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/conf"
	"github.com/RaymondCode/simple-demo/controller/vo"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
)

type VideoListResponse struct {
	vo.Response
	VideoList []vo.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	//鉴权
	if err := middleware.ParseToken(token); err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	user, err := service.GetUserByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	//获取视频数据
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", user.ID, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//保存视频到cdn
	err = service.Upload(finalName)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//将视频信息插入数据库
	err = service.CreateVideo(*user, conf.CDN.Url+finalName, "")
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, vo.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
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
	_, err = service.GetUserByToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid token",
		})
		return
	}
	result, err := service.GetPublishList(userId)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: vo.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
	}
	for i, video := range result {
		result[i].IsFavorite = service.IsFavorite(userId, video.Id)
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: vo.Response{
			StatusCode: 0,
		},
		VideoList: result,
	})
}
