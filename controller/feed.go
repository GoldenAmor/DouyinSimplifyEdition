package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/view_model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []view_model.Video `json:"video_list,omitempty"`
	NextTime  int64              `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	//从数据库中根据用户信息获取发布时间倒序的视频流
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
	result := dao.GetVideos()

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: result,
		NextTime:  time.Now().Unix(),
	})
}
