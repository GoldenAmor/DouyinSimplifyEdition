package rpc

import (
	"context"
	"dousheng/kitex_gen/video"
	"dousheng/kitex_gen/video/videoservice"
	"dousheng/pkg/conf"
	"errors"
	"github.com/cloudwego/kitex/client"
	"time"
)

var videoClient videoservice.Client

func initVideoRpc() {
	c, err := videoservice.NewClient("example", client.WithHostPorts(conf.VideoHostPorts))
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func CreateVideo(ctx context.Context, createVideoReq *video.CreateVideoReq) error {
	resp, err := videoClient.CreateVideo(ctx, createVideoReq)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errors.New(resp.BaseResp.StatusMessage)
	}
	return nil
}

func GetPublishList(ctx context.Context, getPublishListReq *video.GetPublishListReq) ([]*video.Video, error) {
	resp, err := videoClient.GetPublishList(ctx, getPublishListReq)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errors.New(resp.BaseResp.StatusMessage)
	}
	return resp.Videos, nil
}

func GetVideos(ctx context.Context, getVideosReq *video.GetVideosReq) ([]*video.Video, int64, error) {
	resp, err := videoClient.GetVideos(ctx, getVideosReq)
	if err != nil {
		return nil, time.Now().Unix(), err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, time.Now().Unix(), errors.New(resp.BaseResp.StatusMessage)
	}
	return resp.Videos, resp.NextTime, nil
}

func GetVideoById(ctx context.Context, getVideoByIdReq *video.GetVideoByIdReq) (*video.Video, error) {
	resp, err := videoClient.GetVideoById(ctx, getVideoByIdReq)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errors.New(resp.BaseResp.StatusMessage)
	}
	return resp.Video, nil
}
