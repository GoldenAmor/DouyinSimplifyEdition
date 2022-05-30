package dto

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	AuthorId      uint
	Author        User
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
}
