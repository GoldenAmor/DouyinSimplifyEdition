package controller

import (
	"github.com/RaymondCode/simple-demo/controller/vo"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	vo.Response
	VideoList []vo.Video `json:"video_list,omitempty"`
	NextTime  int64      `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	//从数据库中根据用户信息获取发布时间倒序的视频流
	result, err := service.GetVideos()
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: vo.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
	}
	//登陆状态
	token := c.Query("token")
	user, err := service.GetUserByToken(token)
	if err == nil {
		for i, video := range result {
			result[i].IsFavorite = service.IsFavorite(user.ID, video.Id)
		}
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  vo.Response{StatusCode: 0},
		VideoList: result,
		NextTime:  time.Now().Unix(),
	})
}
