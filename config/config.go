package config

type Database struct {
	DSN string
}

type CDN struct {
	AccessKey string
	SecretKey string
	Bucket    string
	CDNUrl    string
}

// EnvVersion 开发版本: develop(开发版) release(正式版)
var EnvVersion = "develop"

// DBDevelop 开发版数据库配置
var DBDevelop = Database{
	DSN: "root:3333@tcp(127.0.0.1:3306)/dousheng?charset=utf8mb4&parseTime=True&loc=Local",
}

// QiNiuCDN 七牛云CDN配置
var QiNiuCDN = CDN{
	AccessKey: "SHFEDQXdwn4MO7-4OvW92Qr0LXzEeR-b7si5gpWW",
	SecretKey: "cbCu0I13O2fLIQq8DbLnhUNydDvNdyaSiAJSwnJj",
	Bucket:    "dousheng2022",
	CDNUrl:    "http://rcnh4fg5p.hn-bkt.clouddn.com/",
}
