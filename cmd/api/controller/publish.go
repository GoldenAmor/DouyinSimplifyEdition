package controller

import (
	"context"
	"dousheng/cmd/api/controller/vo"
	"dousheng/cmd/api/rpc"
	"dousheng/cmd/api/utils"
	"dousheng/kitex_gen/video"
	"dousheng/pkg/middleware"
	"fmt"
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
	userId, err := middleware.GetUserIdByToken(token)
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
	finalName := fmt.Sprintf("%d_%s", userId, filename)
	saveFile := filepath.Join("/public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	go utils.PublishGoroutine(finalName, userId)
	c.JSON(http.StatusOK, vo.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	userIdData := c.Query("user_id")
	targetUserId, err := strconv.ParseInt(userIdData, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid user_id",
		})
		return
	}
	userId, err := middleware.GetUserIdByToken(token)
	if err != nil {
		userId = -1
	}
	videos, err := rpc.GetPublishList(context.Background(), &video.GetPublishListReq{UserId: targetUserId})
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: vo.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
	}
	result := make([]vo.Video, len(videos))
	for i, v := range videos {
		vi, err := utils.PackageVideo(userId, v)
		if err != nil {
			c.JSON(http.StatusOK, VideoListResponse{
				Response: vo.Response{
					StatusCode: 1,
					StatusMsg:  err.Error(),
				},
			})
		}
		result[i] = *vi
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: vo.Response{
			StatusCode: 0,
		},
		VideoList: result,
	})
}
