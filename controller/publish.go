package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/dto"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/view_model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"path/filepath"
)

type VideoListResponse struct {
	Response
	VideoList []view_model.Video `json:"video_list"`
}

// Publish 等待整合鉴权，保存封面未完善（封面数据未给，考虑自动生成封面）
// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//保存视频到cdn
	err = service.Upload(finalName)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//获取videoDao
	dao, err := service.GetVideoInstance()
	if err != nil {
		fmt.Printf(err.Error())
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	//插入视频信息
	dao.CreateVideo(&dto.Video{
		Author: dto.User{
			Model: gorm.Model{
				ID: uint(user.Id),
			},
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
		},
		PlayUrl:       config.QiNiuCDN.CDNUrl + finalName,
		CoverUrl:      "",
		FavoriteCount: 0,
		CommentCount:  0,
	})
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	fmt.Println(token)
	userId := c.Query("user_id")
	dao, err := service.GetVideoInstance()
	if err != nil {
		fmt.Printf(err.Error())
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	result := dao.GetPublishList(userId)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: result,
	})
}
