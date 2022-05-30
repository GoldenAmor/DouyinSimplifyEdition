package service

import (
	"github.com/RaymondCode/simple-demo/dto"
	"github.com/RaymondCode/simple-demo/view_model"
	"sync"
)

type Video struct {
}

var (
	video     *Video
	videoOnce sync.Once
)

func GetVideoInstance() (*Video, error) {
	var err error
	videoOnce.Do(func() {
		err = InitDB()
	})
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (v *Video) CreateVideo(video *dto.Video) error {
	db.Create(video)
	return nil
}

func (v *Video) GetVideos() []view_model.Video {
	var videos []dto.Video
	var result []view_model.Video
	db.Preload("Author").Limit(30).Order("created_at desc").Find(&videos).Scan(&result)
	for i, d := range videos {
		result[i].Author = view_model.User{
			Id:            int64(d.Author.ID),
			Name:          d.Author.Name,
			FollowCount:   d.Author.FollowCount,
			FollowerCount: d.Author.FollowerCount,
			IsFollow:      false,
		}
	}
	return result
}

func (v *Video) GetPublishList(authorId string) []view_model.Video {
	var videos []dto.Video
	var result []view_model.Video
	db.Preload("Author").Limit(30).Order("created_at desc").Where("author_id=?", authorId).Find(&videos).Scan(&result)
	for i, d := range videos {
		result[i].Author = view_model.User{
			Id:            int64(d.Author.ID),
			Name:          d.Author.Name,
			FollowCount:   d.Author.FollowCount,
			FollowerCount: d.Author.FollowerCount,
			IsFollow:      false,
		}
	}
	return result
}
