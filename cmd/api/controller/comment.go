package controller

import (
	"context"
	"dousheng/cmd/api/controller/vo"
	"dousheng/cmd/api/rpc"
	"dousheng/kitex_gen/comment"
	"dousheng/kitex_gen/relation"
	"dousheng/pkg/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	vo.Response
	CommentList []vo.Comment `json:"comment_list,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
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
	//评论
	if actionType == "1" {
		content := c.Query("comment_text")
		err = rpc.CreateComment(context.Background(), &comment.CreateCommentReq{
			UserId:  userId,
			VideoId: videoId,
			Content: content,
		})
	}
	//删除评论
	if actionType == "2" {
		commentIdData := c.Query("comment_id")
		commentId, err := strconv.ParseInt(commentIdData, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "comment_id error"})
			return
		}
		err = rpc.DeleteComment(context.Background(), &comment.DeleteCommentReq{
			VideoId:   videoId,
			CommentId: commentId,
		})
	}
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, vo.Response{StatusCode: 0})
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	token := c.Query("token")
	//鉴权
	userId, err := middleware.GetUserIdByToken(token)
	if err != nil {
		userId = -1
	}
	videoIdData := c.Query("video_id")
	videoId, err := strconv.ParseInt(videoIdData, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "video_id error"})
		return
	}
	comments, err := rpc.GetCommentsByVideoId(context.Background(), &comment.GetCommentsByVideoIdReq{VideoId: videoId})
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	result := make([]vo.Comment, len(comments))
	for i, co := range comments {
		isFollow, err := rpc.IsFollow(context.Background(), &relation.IsFollowReq{
			UserId:       userId,
			TargetUserId: co.User.Id,
		})
		if err != nil {
			c.JSON(http.StatusOK, vo.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
		}
		followCount, err := rpc.CountFollows(context.Background(), &relation.CountFollowsReq{UserId: co.User.Id})
		if err != nil {
			c.JSON(http.StatusOK, vo.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
		}
		followerCount, err := rpc.CountFollowers(context.Background(), &relation.CountFollowersReq{UserId: co.User.Id})
		if err != nil {
			c.JSON(http.StatusOK, vo.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
		}
		result[i] = vo.Comment{
			Id: co.Id,
			User: vo.User{
				Id:            co.User.Id,
				Name:          co.User.Name,
				FollowCount:   followCount,
				FollowerCount: followerCount,
				IsFollow:      isFollow,
			},
			Content:    co.Content,
			CreateDate: co.CreateDate,
		}
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    vo.Response{StatusCode: 0},
		CommentList: result,
	})
}
