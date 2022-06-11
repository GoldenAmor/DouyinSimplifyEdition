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
	UserId    int64
	ByUserId  int64
}
