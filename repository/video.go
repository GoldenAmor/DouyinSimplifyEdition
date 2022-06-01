package repository

import (
	"gorm.io/gorm"
	"time"
)

type Video struct {
	ID            int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	AuthorId      int64
	Author        User
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
}
