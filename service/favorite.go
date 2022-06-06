package service

import (
	"github.com/RaymondCode/simple-demo/controller/vo"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/repository"
)

func IsFavorite(userId int64, videoId int64) bool {
	favoriteDao := dao.GetFavoriteDaoInstance()
	_, err := favoriteDao.GetByUserIdAndVideoId(userId, videoId)
	if err != nil {
		return false
	}
	return true
}

func Like(userId int64, videoId int64) error {
	lock.Lock()
	defer lock.Unlock()
	video, err := GetVideoById(videoId)
	if err != nil {
		return err
	}
	//数据库保存点赞操作
	favoriteDao := dao.GetFavoriteDaoInstance()
	err = favoriteDao.Create(&repository.Favorite{
		UserId:  userId,
		VideoId: videoId,
	})
	if err != nil {
		return err
	}
	//更新视频点赞数据
	err = UpdateVideoFavoriteCount(videoId, video.FavoriteCount+1)
	if err != nil {
		return err
	}
	return nil
}

func UnLike(userId int64, videoId int64) error {
	lock.Lock()
	defer lock.Unlock()
	video, err := GetVideoById(videoId)
	//数据库保存取消点赞操作
	if err != nil {
		return err
	}
	favoriteDao := dao.GetFavoriteDaoInstance()
	target, err := favoriteDao.GetByUserIdAndVideoId(userId, videoId)
	if err != nil {
		return err
	}
	err = favoriteDao.Delete(target)
	if err != nil {
		return err
	}
	//更新视频点赞数据
	err = UpdateVideoFavoriteCount(videoId, video.FavoriteCount-1)
	if err != nil {
		return err
	}
	return nil
}

func GetFavoritesByUserId(userId int64) ([]vo.Video, error) {
	favoriteDao := dao.GetFavoriteDaoInstance()
	videoDao := dao.GetVideoInstance()
	favorites, err := favoriteDao.GetByUserId(userId)
	if err != nil {
		return nil, err
	}
	result := make([]vo.Video, len(favorites))
	for i, favorite := range favorites {
		v, err := videoDao.GetVideoById(favorite.VideoId)
		if err != nil {
			return nil, err
		}
		result[i] = Transform2VoVideo(*v)
	}
	return result, nil
}
