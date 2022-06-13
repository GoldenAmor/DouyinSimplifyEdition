package dao

import (
	"dousheng/cmd/video/conn"
	"dousheng/cmd/video/repository"
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

func (v *Video) GetVideos(latestTime string) ([]repository.Video, error) {
	var videos []repository.Video
	err := conn.DB.Preload("Author").Limit(2).Order("created_at desc").Where("created_at < ?", latestTime).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (v *Video) GetVideoById(id int64) (*repository.Video, error) {
	result := &repository.Video{}
	err := conn.DB.Preload("Author").Where("id = ?", id).First(result).Error
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
