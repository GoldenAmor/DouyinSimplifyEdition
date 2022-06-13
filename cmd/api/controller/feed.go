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
	"strconv"
	"time"
)

type FeedResponse struct {
	vo.Response
	VideoList []vo.Video `json:"video_list,omitempty"`
	NextTime  int64      `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	latestTimeData := c.Query("latest_time")
	latestTime, err := strconv.ParseInt(latestTimeData, 10, 64)

	if err != nil {
		fmt.Println(err.Error())
		latestTime = time.Now().Unix()
	}
	if latestTime > 10000000000 {
		latestTime = latestTime / 1000
	}
	fmt.Println(latestTime)
	//登陆状态
	token := c.Query("token")
	userId, err := middleware.GetUserIdByToken(token)
	if err != nil {
		userId = -1
	}
	//从数据库中根据用户信息获取发布时间倒序的视频流
	videos, nextTime, err := rpc.GetVideos(context.Background(), &video.GetVideosReq{
		LatestTime: latestTime,
	})
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
	fmt.Printf("%#v\n", result)
	c.JSON(http.StatusOK, FeedResponse{
		Response:  vo.Response{StatusCode: 0},
		VideoList: result,
		NextTime:  nextTime,
	})
}
