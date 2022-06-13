package service

import (
	"context"
	"dousheng/cmd/video/dao"
	"dousheng/kitex_gen/video"
	"time"
)

type QueryVideoServiceImpl struct {
	ctx context.Context
}

func NewQueryVideoServiceImpl(ctx context.Context) *QueryVideoServiceImpl {
	return &QueryVideoServiceImpl{ctx: ctx}
}

func (receiver *QueryVideoServiceImpl) GetPublishList(userId int64) ([]*video.Video, error) {
	videoDao := dao.GetVideoInstance()
	videos, err := videoDao.GetVideosByAuthorId(userId)
	if err != nil {
		return nil, err
	}
	result := make([]*video.Video, len(videos))
	for i, v := range videos {
		result[i] = &video.Video{
			Id:       v.ID,
			AuthorId: v.AuthorId,
			Author: &video.User{
				Id:   v.Author.ID,
				Name: v.Author.Name,
			},
			PlayUrl:  v.PlayUrl,
			CoverUrl: v.CoverUrl,
		}
	}
	return result, nil
}

func (receiver *QueryVideoServiceImpl) GetVideos(latestTime int64) ([]*video.Video, int64, error) {
	videoDao := dao.GetVideoInstance()
	var timeLayoutStr = "2006-01-02 15:04:05"
	videos, err := videoDao.GetVideos(time.Unix(latestTime, 0).Format(timeLayoutStr))
	if err != nil {
		return nil, time.Now().Unix(), err
	}
	result := make([]*video.Video, len(videos))
	for i, v := range videos {
		result[i] = &video.Video{
			Id:       v.ID,
			AuthorId: v.AuthorId,
			Author: &video.User{
				Id:   v.Author.ID,
				Name: v.Author.Name,
			},
			PlayUrl:  v.PlayUrl,
			CoverUrl: v.CoverUrl,
		}
	}
	if len(result) == 0 {
		return result, time.Now().Unix(), nil
	}
	return result, videos[len(videos)-1].CreatedAt.Unix(), nil
}

func (receiver *QueryVideoServiceImpl) GetVideoById(videoId int64) (*video.Video, error) {
	videoDao := dao.GetVideoInstance()
	v, err := videoDao.GetVideoById(videoId)
	if err != nil {
		return nil, err
	}
	return &video.Video{
		Id:       v.ID,
		AuthorId: v.AuthorId,
		Author: &video.User{
			Id:   v.Author.ID,
			Name: v.Author.Name,
		},
		PlayUrl:  v.PlayUrl,
		CoverUrl: v.CoverUrl,
	}, nil
}
