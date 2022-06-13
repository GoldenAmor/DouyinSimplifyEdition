package service

import (
	"context"
	"dousheng/cmd/video/dao"
	"dousheng/cmd/video/repository"
)

type CreateVideoServiceImpl struct {
	ctx context.Context
}

func NewCreateVideoServiceImpl(ctx context.Context) *CreateVideoServiceImpl {
	return &CreateVideoServiceImpl{ctx: ctx}
}

func (receiver CreateVideoServiceImpl) CreateVideo(userId int64, playUrl string, coverUrl string) error {
	videoDao := dao.GetVideoInstance()
	return videoDao.Create(repository.Video{
		AuthorId: userId,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
	})
}
