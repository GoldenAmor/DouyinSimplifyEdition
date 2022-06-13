package rpc

import (
	"context"
	"dousheng/kitex_gen/comment"
	"dousheng/kitex_gen/comment/commentservice"
	"dousheng/pkg/conf"
	"errors"
	"github.com/cloudwego/kitex/client"
)

var commentClient commentservice.Client

func initCommentRpc() {
	c, err := commentservice.NewClient("example", client.WithHostPorts(conf.CommentHostPorts))
	if err != nil {
		panic(err)
	}
	commentClient = c
}

func CreateComment(ctx context.Context, createCommentReq *comment.CreateCommentReq) error {
	resp, err := commentClient.CreateComment(ctx, createCommentReq)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errors.New(resp.BaseResp.StatusMessage)
	}
	return nil
}

// DeleteComment implements the CommentServiceImpl interface.
func DeleteComment(ctx context.Context, deleteCommentReq *comment.DeleteCommentReq) error {
	resp, err := commentClient.DeleteComment(ctx, deleteCommentReq)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errors.New(resp.BaseResp.StatusMessage)
	}
	return nil
}

// GetCommentsByVideoId implements the CommentServiceImpl interface.
func GetCommentsByVideoId(ctx context.Context, getCommentsByVideoIdReq *comment.GetCommentsByVideoIdReq) ([]*comment.Comment, error) {
	resp, err := commentClient.GetCommentsByVideoId(ctx, getCommentsByVideoIdReq)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errors.New(resp.BaseResp.StatusMessage)
	}
	return resp.Comments, nil
}

// CountComment implements the CommentServiceImpl interface.
func CountComment(ctx context.Context, countCommentReq *comment.CountCommentReq) (int64, error) {
	// TODO: Your code here...
	resp, err := commentClient.CountComment(ctx, countCommentReq)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errors.New(resp.BaseResp.StatusMessage)
	}
	return resp.CommentCount, nil
}
