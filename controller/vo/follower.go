package vo

type Follower struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	CreateDate string `json:"create_date,omitempty"`
}
