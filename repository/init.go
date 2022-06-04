package repository

import "github.com/RaymondCode/simple-demo/conn"

// Init 创建数据库表
func Init() error {
	err := conn.InitGorm()
	if err != nil {
		return err
	}
	err = conn.DB.AutoMigrate(&Video{}, &User{})
	if err != nil {
		return err
	}
	return nil
}
