package common

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/controller"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBconnection() (*gorm.DB, error) {
	dbConnect := controller.DBConnect{
		"root",
		"123456",
		"43.138.135.43",
		3306,
		"Dousheng",
		"10s",
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", dbConnect.Username, dbConnect.Password, dbConnect.Host, dbConnect.Port, dbConnect.DbName, dbConnect.Timeout)
	db, connectErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if connectErr != nil {
		return nil, connectErr
	} else {
		return db, nil
	}
}
