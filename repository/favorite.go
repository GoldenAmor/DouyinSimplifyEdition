package repository

type Favorite struct {
	Id      int64
	UserId  int64
	User    User
	VideoId int64
	Video   Video
}
