package service

import (
	"context"
	"dousheng/cmd/comment/dao"
	"dousheng/kitex_gen/comment"
)

type QueryCommentServiceImpl struct {
	ctx context.Context
}

func NewQueryCommentServiceImpl(ctx context.Context) *QueryCommentServiceImpl {
	return &QueryCommentServiceImpl{ctx: ctx}
}

func (receiver *QueryCommentServiceImpl) GetCommentsByVideoId(videoId int64) ([]*comment.Comment, error) {
	commentDao := dao.GetCommentDaoInstance()
	comments, err := commentDao.GetListByVideoId(videoId)
	if err != nil {
		return nil, err
	}
	var timeLayoutStr = "01-02"
	result := make([]*comment.Comment, len(comments))
	for i, c := range comments {
		result[i] = &comment.Comment{
			Id: c.Id,
			User: &comment.User{
				Id:       c.User.ID,
				Name:     c.User.Name,
				Password: c.User.Password,
			},
			Content:    c.Content,
			CreateDate: c.CreatedAt.Format(timeLayoutStr),
		}
	}
	return result, nil
}
