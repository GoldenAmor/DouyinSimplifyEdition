package dao

import (
	"github.com/RaymondCode/simple-demo/conn"
	"github.com/RaymondCode/simple-demo/repository"
	"sync"
)

type Video struct {
}

var (
	video     *Video
	videoOnce sync.Once
)

func GetVideoInstance() *Video {
	videoOnce.Do(func() {
		video = &Video{}
	})
	return video
}

func (v *Video) Create(object repository.Video) error {
	err := conn.DB.Create(&object).Error
	if err != nil {
		return err
	}
	return nil
}

func (v *Video) GetVideos() ([]repository.Video, error) {
	var videos []repository.Video
	err := conn.DB.Preload("Author").Limit(30).Order("created_at desc").Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (v *Video) GetVideoById(id int64) (*repository.Video, error) {
	result := &repository.Video{}
	err := conn.DB.Where("id = ?", id).First(result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (v *Video) GetVideosByAuthorId(authorId int64) ([]repository.Video, error) {
	var videos []repository.Video
	err := conn.DB.Preload("Author").Limit(30).Order("created_at desc").Where("author_id=?", authorId).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (v *Video) UpdateVideoFavorite(videoId int64, count int64) error {
	return conn.DB.Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", count).Error
}

func (v *Video) UpdateVideoCommentCount(videoId int64, count int64) error {
	return conn.DB.Model(&Video{}).Where("id = ?", videoId).Update("comment_count", count).Error
}
