package repository

import (
	"gorm.io/gorm"
	"time"
)

type Follow struct {
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserId    int64
	ToUserId  int64
}
