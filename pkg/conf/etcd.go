package conf

const (
	UserPort          = ":8801"
	VideoPort         = ":8802"
	FavoritePort      = ":8803"
	CommentPort       = ":8804"
	RelationPort      = ":8805"
	UserHostPorts     = "0.0.0.0" + UserPort
	VideoHostPorts    = "0.0.0.0" + VideoPort
	FavoriteHostPorts = "0.0.0.0" + FavoritePort
	CommentHostPorts  = "0.0.0.0" + CommentPort
	RelationHostPorts = "0.0.0.0" + RelationPort
)
