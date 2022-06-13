package dao

import (
	"dousheng/cmd/favorite/conn"
	"dousheng/cmd/favorite/repository"
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

func (f *Favorite) CreateFavorite(object *repository.Favorite) error {
	return conn.DB.Create(object).Error
}

func (f *Favorite) DeleteFavorite(object *repository.Favorite) error {
	return conn.DB.Delete(object).Error
}

func (f *Favorite) GetFavoriteByUserIdAndVideoId(userId int64, videoId int64) (*repository.Favorite, error) {
	result := &repository.Favorite{}
	err := conn.DB.Where("user_id = ?", userId).Where("video_id = ?", videoId).First(result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (f *Favorite) GetFavoritesByUserId(userId int64) ([]repository.Favorite, error) {
	var favorites []repository.Favorite
	err := conn.DB.Where("user_id = ?", userId).Find(&favorites).Error
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

func (f *Favorite) ContainsFavorite(userId int64, videoId int64) (bool, error) {
	var count int64
	err := conn.DB.Model(&repository.Favorite{}).Where("user_id = ?", userId).Where("video_id = ?", videoId).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count != 0, nil
}

func (f *Favorite) CountFavorite(videoId int64) (int64, error) {
	var count int64
	err := conn.DB.Model(&repository.Favorite{}).Where("video_id = ?", videoId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
