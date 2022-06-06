package dao

import (
	"github.com/RaymondCode/simple-demo/conn"
	"github.com/RaymondCode/simple-demo/repository"
	"sync"
)

type Favorite struct{}

var (
	favorite     Favorite
	favoriteOnce sync.Once
)

func GetFavoriteDaoInstance() *Favorite {
	favoriteOnce.Do(func() {
		favorite = Favorite{}
	})
	return &favorite
}

func (f *Favorite) Create(object *repository.Favorite) error {
	return conn.DB.Create(object).Error
}

func (f *Favorite) Delete(object *repository.Favorite) error {
	return conn.DB.Delete(object).Error
}

func (f *Favorite) GetByUserIdAndVideoId(userId int64, videoId int64) (*repository.Favorite, error) {
	result := &repository.Favorite{}
	err := conn.DB.Where("user_id = ?", userId).Where("video_id = ?", videoId).First(result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (f *Favorite) GetByUserId(userId int64) ([]repository.Favorite, error) {
	var favorites []repository.Favorite
	err := conn.DB.Where("user_id = ?", userId).Find(&favorites).Error
	if err != nil {
		return nil, err
	}
	return favorites, nil
}
