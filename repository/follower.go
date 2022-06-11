package repository

import (
	"gorm.io/gorm"
	"time"
)

type Follower struct {
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	User      User
	UserId    int64
	ByUser    User
	ByUserId  int64
}
