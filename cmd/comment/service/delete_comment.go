package service

import (
	"context"
	"dousheng/cmd/comment/dao"
	"dousheng/cmd/comment/repository"
)

type DeleteCommentServiceImpl struct {
	ctx context.Context
}

func NewDeleteCommentServiceImpl(ctx context.Context) *DeleteCommentServiceImpl {
	return &DeleteCommentServiceImpl{ctx: ctx}
}

func (receiver *DeleteCommentServiceImpl) DeleteComment(commentId int64) error {
	commentDao := dao.GetCommentDaoInstance()
	//删除评论
	return commentDao.Delete(repository.Comment{
		Id: commentId,
	})
}
