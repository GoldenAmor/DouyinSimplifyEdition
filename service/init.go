package service

import (
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/dto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	db       *gorm.DB
	initOnce sync.Once
)

// InitDB 获取db连接
func InitDB() error {
	var err error
	//初始化 只执行一次
	initOnce.Do(func() {
		db, err = gorm.Open(mysql.Open(config.DBDevelop.DSN), &gorm.Config{})
	})
	if err != nil {
		return err
	}
	return nil
}

// Init 创建数据库表
func Init() error {
	err := InitDB()
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&dto.Video{}, &dto.User{})
	if err != nil {
		return err
	}
	return nil
}
