package controller

import (
	"github.com/RaymondCode/simple-demo/controller/vo"
	"github.com/RaymondCode/simple-demo/public"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommentListResponse struct {
	vo.Response
	CommentList []vo.Comment `json:"comment_list,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := public.UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, vo.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    vo.Response{StatusCode: 0},
		CommentList: public.DemoComments,
	})
}
