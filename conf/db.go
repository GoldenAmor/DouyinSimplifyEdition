package conf

type db struct {
	Username string
	Password string
	Host     string
	Port     int
	DbName   string
	Timeout  string
}

var MysqlDB = db{
	Username: "root",
	Password: "123456",
	Host:     "43.138.135.43",
	Port:     3306,
	DbName:   "Dousheng",
	Timeout:  "10s",
}
