package dto

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name          string
	FollowCount   int64
	FollowerCount int64
}
