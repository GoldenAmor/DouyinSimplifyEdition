package repository

type Relation struct {
	Id         int64
	UserId     int64
	User       User
	FollowerId int64
	Follower   User
}
