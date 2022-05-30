/*****************************************************************************
Author:Wangfeng
Date:May 24
Introduction:实现与远程Mysql服务器的连接。
 *****************************************************************************/
package main

import(
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct{
	userpassword 		string
	Id			 		int
	Name		 		string
	FollowCount	 		int
	FollowerCount 		int
	IsFollow			bool
}

func (u User) TableName() string{
	return "user_info"
}

func ConnectDB() {
	username:="root"		//Mysql测试用户，只有连接权限
	password:="123456"
	host:="43.138.135.43"	//Wangfeng的腾讯云服务器的公网IP
	port:=3306
	DBname:="Dousheng"
	timeout:="10s"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, DBname,timeout)

	db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{})

	if err!=nil{
		panic("连接数据库失败，error="+err.Error())
	}else{
		fmt.Println("连接数据库成功")
	}
	u:=User{}
	db.Take(&u)
	fmt.Println(u)
}

