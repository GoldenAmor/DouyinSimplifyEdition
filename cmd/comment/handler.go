package main

import (
	"context"
	"dousheng/cmd/comment/service"
	"dousheng/kitex_gen/comment"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CreateComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CreateComment(ctx context.Context, createCommentReq *comment.CreateCommentReq) (resp *comment.CreateCommentResp, err error) {
	// TODO: Your code here...
	videoId := createCommentReq.VideoId
	userId := createCommentReq.UserId
	content := createCommentReq.Content
	createCommentService := service.NewCreateCommentServiceImpl(ctx)
	err = createCommentService.CreateComment(userId, videoId, content)
	if err != nil {
		resp = &comment.CreateCommentResp{BaseResp: &comment.BaseResp{
			StatusCode:    1,
			StatusMessage: err.Error(),
		}}
	}
	resp = &comment.CreateCommentResp{BaseResp: &comment.BaseResp{
		StatusCode: 0,
	}}
	return
}

// DeleteComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) DeleteComment(ctx context.Context, deleteCommentReq *comment.DeleteCommentReq) (resp *comment.DeleteCommentResp, err error) {
	// TODO: Your code here...
	//videoId := deleteCommentReq.VideoId
	commentId := deleteCommentReq.CommentId
	createCommentService := service.NewDeleteCommentServiceImpl(ctx)
	err = createCommentService.DeleteComment(commentId)
	if err != nil {
		resp = &comment.DeleteCommentResp{BaseResp: &comment.BaseResp{
			StatusCode:    1,
			StatusMessage: err.Error(),
		}}
	}
	resp = &comment.DeleteCommentResp{BaseResp: &comment.BaseResp{
		StatusCode: 0,
	}}
	return
}

// GetCommentsByVideoId implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) GetCommentsByVideoId(ctx context.Context, getCommentsByVideoIdReq *comment.GetCommentsByVideoIdReq) (resp *comment.GetCommentsByVideoIdResp, err error) {
	// TODO: Your code here...
	videoId := getCommentsByVideoIdReq.VideoId
	queryCommentService := service.NewQueryCommentServiceImpl(ctx)
	comments, err := queryCommentService.GetCommentsByVideoId(videoId)
	if err != nil {
		resp = &comment.GetCommentsByVideoIdResp{
			BaseResp: &comment.BaseResp{
				StatusCode:    1,
				StatusMessage: err.Error(),
			},
			Comments: nil,
		}
		return
	}
	resp = &comment.GetCommentsByVideoIdResp{
		BaseResp: &comment.BaseResp{
			StatusCode: 0,
		},
		Comments: comments,
	}
	return
}

// CountComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CountComment(ctx context.Context, countCommentReq *comment.CountCommentReq) (resp *comment.CountCommentResp, err error) {
	// TODO: Your code here...
	videoId := countCommentReq.VideoId
	queryCommentService := service.NewQueryCommentServiceImpl(ctx)
	comments, err := queryCommentService.GetCommentsByVideoId(videoId)
	if err != nil {
		resp = &comment.CountCommentResp{
			BaseResp: &comment.BaseResp{
				StatusCode:    1,
				StatusMessage: err.Error(),
			},
		}
		return
	}
	resp = &comment.CountCommentResp{
		BaseResp: &comment.BaseResp{
			StatusCode: 0,
		},
		CommentCount: int64(len(comments)),
	}
	return
}
