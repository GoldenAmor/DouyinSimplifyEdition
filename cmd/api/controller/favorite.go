package controller

import (
	"context"
	"dousheng/cmd/api/controller/vo"
	"dousheng/cmd/api/rpc"
	"dousheng/cmd/api/utils"
	"dousheng/kitex_gen/favorite"
	"dousheng/kitex_gen/video"
	"dousheng/pkg/middleware"
	"fmt"
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
	userId, err := middleware.GetUserIdByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	//点赞
	if actionType == "1" {
		err = rpc.Like(context.Background(), &favorite.LikeReq{
			UserId:  userId,
			VideoId: videoId,
		})
	}
	//取消点赞
	if actionType == "2" {
		err = rpc.UnLike(context.Background(), &favorite.UnLikeReq{
			UserId:  userId,
			VideoId: videoId,
		})
	}
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, vo.Response{StatusCode: 0})
}

// FavoriteList all users have same favorite.thrift video list
func FavoriteList(c *gin.Context) {
	//var DemoUser = vo.User{
	//	Id:            1,
	//	Name:          "TestUser",
	//	FollowCount:   0,
	//	FollowerCount: 0,
	//	IsFollow:      false,
	//}
	//var DemoVideos = []vo.Video{
	//	{
	//		Id:            1,
	//		Author:        DemoUser,
	//		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
	//		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
	//		FavoriteCount: 0,
	//		CommentCount:  0,
	//		IsFavorite:    false,
	//	},
	//}
	//c.JSON(http.StatusOK, VideoListResponse{
	//	Response: vo.Response{
	//		StatusCode: 0,
	//	},
	//	VideoList: DemoVideos,
	//})
	token := c.Query("token")
	userIdData := c.Query("user_id")
	targetUserId, err := strconv.ParseInt(userIdData, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid user_id",
		})
		return
	}
	//鉴权
	userId, err := middleware.GetUserIdByToken(token)
	if err != nil {
		userId = -1
	}
	//获取列表
	favorites, err := rpc.GetFavoritesByUserId(context.Background(), &favorite.GetFavoritesByUserIdReq{UserId: targetUserId})
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	result := make([]vo.Video, len(favorites))
	for i, f := range favorites {
		v, err := rpc.GetVideoById(context.Background(), &video.GetVideoByIdReq{Id: f})
		if err != nil {
			c.JSON(http.StatusOK, VideoListResponse{
				Response: vo.Response{
					StatusCode: 1,
					StatusMsg:  err.Error(),
				},
			})
		}
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
	c.JSON(http.StatusOK, VideoListResponse{
		Response: vo.Response{
			StatusCode: 0,
		},
		VideoList: result,
	})
}
