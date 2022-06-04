package service

import (
	"github.com/RaymondCode/simple-demo/controller/vo"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/repository"
)

func CreateVideo(user repository.User, playUrl string, coverUrl string) error {
	videoDao := dao.GetVideoInstance()
	return videoDao.Create(&repository.Video{
		AuthorId:      user.ID,
		PlayUrl:       playUrl,
		CoverUrl:      coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
	})
}
func GetPublishList(userId string) ([]vo.Video, error) {
	videoDao := dao.GetVideoInstance()
	videos, err := videoDao.GetVideosByAuthorId(userId)
	if err != nil {
		return nil, err
	}
	result := make([]vo.Video, len(videos))
	for i, video := range videos {
		result[i] = Transform2VoVideo(video)
	}
	return result, nil
}

func GetVideos() ([]vo.Video, error) {
	videoDao := dao.GetVideoInstance()
	videos, err := videoDao.GetVideos()
	if err != nil {
		return nil, err
	}
	result := make([]vo.Video, len(videos))
	for i, video := range videos {
		result[i] = Transform2VoVideo(video)
	}
	return result, nil
}

func Transform2VoVideo(video repository.Video) vo.Video {
	return vo.Video{
		Id: video.ID,
		Author: vo.User{
			Id:            video.Author.ID,
			Name:          video.Author.Name,
			FollowCount:   video.Author.FollowCount,
			FollowerCount: video.Author.FollowerCount,
			IsFollow:      false,
		},
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    false,
	}
}
