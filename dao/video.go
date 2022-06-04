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

func (v *Video) Create(video *repository.Video) error {
	err := conn.DB.Create(video).Error
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

func (v *Video) GetVideosByAuthorId(authorId string) ([]repository.Video, error) {
	var videos []repository.Video
	err := conn.DB.Preload("Author").Limit(30).Order("created_at desc").Where("author_id=?", authorId).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}
