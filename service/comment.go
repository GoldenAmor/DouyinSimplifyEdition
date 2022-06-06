package service

import (
	"github.com/RaymondCode/simple-demo/controller/vo"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/repository"
)

func CreateComment(userId int64, videoId int64, content string) error {
	lock.Lock()
	defer lock.Unlock()
	video, err := GetVideoById(videoId)
	if err != nil {
		return err
	}
	commentDao := dao.GetCommentDaoInstance()
	//保存评论
	err = commentDao.Create(repository.Comment{
		UserId:  userId,
		VideoId: videoId,
		Content: content,
	})
	if err != nil {
		return err
	}
	//更新视频数据
	err = UpdateVideoCommentCount(videoId, video.FavoriteCount+1)
	if err != nil {
		return err
	}
	return nil
}

func DeleteComment(commentId int64, videoId int64) error {
	lock.Lock()
	defer lock.Unlock()
	video, err := GetVideoById(videoId)
	if err != nil {
		return err
	}
	//删除评论
	commentDao := dao.GetCommentDaoInstance()
	err = commentDao.Delete(repository.Comment{
		Id: commentId,
	})
	if err != nil {
		return err
	}
	//更新视频数据
	err = UpdateVideoCommentCount(videoId, video.FavoriteCount-1)
	if err != nil {
		return err
	}
	return nil
}

func GetCommentsByVideoId(videoId int64) ([]vo.Comment, error) {
	commentDao := dao.GetCommentDaoInstance()
	comments, err := commentDao.GetListByVideoId(videoId)
	if err != nil {
		return nil, err
	}
	var timeLayoutStr = "01-02"
	result := make([]vo.Comment, len(comments))
	for i, comment := range comments {
		result[i] = vo.Comment{
			Id:         comment.Id,
			User:       *Transform2VoUser(GetUserById(comment.UserId)),
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.Format(timeLayoutStr),
		}
	}
	return result, nil
}
