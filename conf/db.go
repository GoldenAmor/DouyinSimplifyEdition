package conf

type db struct {
	Username string //Mysql测试用户，只有连接权限
	Password string
	Host     string //Wangfeng的腾讯云服务器的公网IP
	Port     int
	DbName   string
	Timeout  string
}

var DB = db{
	Username: "root",
	Password: "123456",
	Host:     "43.138.135.43",
	Port:     3306,
	DbName:   "Dousheng",
	Timeout:  "10s",
}
