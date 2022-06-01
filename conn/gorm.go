package conn

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitGorm() error {
	dbConnect := conf.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", dbConnect.Username, dbConnect.Password, dbConnect.Host, dbConnect.Port, dbConnect.DbName, dbConnect.Timeout)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
