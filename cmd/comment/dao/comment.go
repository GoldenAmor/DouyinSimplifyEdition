package dao

import (
	"dousheng/cmd/comment/conn"
	"dousheng/cmd/comment/repository"
	"sync"
)

type Comment struct{}

var (
	comment     Comment
	commentOnce sync.Once
)

func GetCommentDaoInstance() *Comment {
	commentOnce.Do(func() {
		comment = Comment{}
	})
	return &comment
}

func (c *Comment) Create(object repository.Comment) error {
	return conn.DB.Create(&object).Error
}

func (c *Comment) Delete(object repository.Comment) error {
	return conn.DB.Delete(&object).Error
}

func (c *Comment) GetListByVideoId(videoId int64) ([]repository.Comment, error) {
	var result []repository.Comment
	err := conn.DB.Preload("User").Order("created_at desc").Where("video_id = ?", videoId).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
