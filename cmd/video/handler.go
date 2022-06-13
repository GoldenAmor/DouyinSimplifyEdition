package main

import (
	"context"
	"dousheng/cmd/video/service"
	"dousheng/kitex_gen/video"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// CreateVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CreateVideo(ctx context.Context, createVideoReq *video.CreateVideoReq) (resp *video.CreateVideoResp, err error) {
	// TODO: Your code here...
	userId := createVideoReq.UserId
	playUrl := createVideoReq.PlayUrl
	coverUrl := createVideoReq.CoverUrl
	createVideoService := service.NewCreateVideoServiceImpl(ctx)
	err = createVideoService.CreateVideo(userId, playUrl, coverUrl)
	if err != nil {
		resp = &video.CreateVideoResp{BaseResp: &video.BaseResp{
			StatusCode:    1,
			StatusMessage: err.Error(),
		}}
		return
	}
	resp = &video.CreateVideoResp{BaseResp: &video.BaseResp{
		StatusCode: 0,
	}}
	return
}

// GetPublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishList(ctx context.Context, getPublishListReq *video.GetPublishListReq) (resp *video.GetPublishListResp, err error) {
	// TODO: Your code here...
	userId := getPublishListReq.UserId
	queryVideoService := service.NewQueryVideoServiceImpl(ctx)
	videos, err := queryVideoService.GetPublishList(userId)
	if err != nil {
		resp = &video.GetPublishListResp{
			BaseResp: &video.BaseResp{
				StatusCode:    1,
				StatusMessage: err.Error(),
			},
		}
		return
	}
	resp = &video.GetPublishListResp{
		BaseResp: &video.BaseResp{
			StatusCode: 0,
		},
		Videos: videos,
	}
	return
}

// GetVideos implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideos(ctx context.Context, getVideosReq *video.GetVideosReq) (resp *video.GetVideosResp, err error) {
	// TODO: Your code here...
	//从数据库中根据用户信息获取发布时间倒序的视频流
	latestTime := getVideosReq.LatestTime
	queryVideoService := service.NewQueryVideoServiceImpl(ctx)
	result, nextTime, err := queryVideoService.GetVideos(latestTime)
	if err != nil {
		resp = &video.GetVideosResp{
			BaseResp: &video.BaseResp{
				StatusCode:    1,
				StatusMessage: err.Error(),
			},
		}
		return
	}
	resp = &video.GetVideosResp{
		BaseResp: &video.BaseResp{
			StatusCode: 0,
		},
		Videos:   result,
		NextTime: nextTime,
	}
	return
}

// GetVideoById implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoById(ctx context.Context, getVideoByIdReq *video.GetVideoByIdReq) (resp *video.GetVideoByIdResp, err error) {
	// TODO: Your code here...
	videoId := getVideoByIdReq.Id
	queryVideoService := service.NewQueryVideoServiceImpl(ctx)
	v, err := queryVideoService.GetVideoById(videoId)
	if err != nil {
		resp = &video.GetVideoByIdResp{
			BaseResp: &video.BaseResp{
				StatusCode:    1,
				StatusMessage: err.Error(),
			},
		}
		return
	}
	resp = &video.GetVideoByIdResp{
		BaseResp: &video.BaseResp{
			StatusCode: 0,
		},
		Video: v,
	}
	return
}

// Publish implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Publish(ctx context.Context, publishReq *video.PublishReq) (resp *video.PublishResp, err error) {
	// TODO: Your code here...
	return
}
