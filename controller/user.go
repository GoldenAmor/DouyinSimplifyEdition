package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

//var dbConnect = DBConnect{
//	"root",
//	"123456",
//	"43.138.135.43",
//	3306,
//	"Dousheng",
//	"10s",
//}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", dbConnect.username, dbConnect.password, dbConnect.host, dbConnect.port, dbConnect.DBname, dbConnect.timeout)
	db, connectErr := common.DBconnection()
	if connectErr != nil {
		fmt.Println("Fail to connect database!")
	}

	u := User{}
	nameErr := db.Table("user_info").Where("name = ?", username).First(&u).Error

	if nameErr == nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "User already existed",
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		newUser := User{
			userIdSequence, username, 0, 0, false,
		}
		CreateErr := db.Table("user_info").Model(&User{}).Omit("Id").Create(&newUser).Error

		if CreateErr != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 2,
				StatusMsg:  "Unexpected fault!",
			})
		} else {
			token, err := common.GenerateToken(username, password)
			if err != nil {
				c.JSON(http.StatusOK, Response{
					StatusCode: 3,
					StatusMsg:  "Fail to generate token",
				})
			} else {
				var Id int64
				db.Table("user_info").Select("id").Where("name = ?", username).Find(&Id)
				c.JSON(http.StatusOK, UserLoginResponse{
					Response: Response{0, "Register success!"},
					UserId:   Id,
					Token:    token,
				})
			}
		}
	}

	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, UserLoginResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
	//	})
	//} else {
	//	atomic.AddInt64(&userIdSequence, 1)
	//	newUser := User{
	//		Id:   userIdSequence,
	//		Name: username,
	//	}
	//	usersLoginInfo[token] = newUser
	//	c.JSON(http.StatusOK, UserLoginResponse{
	//		Response: Response{StatusCode: 0},
	//		UserId:   userIdSequence,
	//		Token:    username + password,
	//	})
	//}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password
	db, connectErr := common.DBconnection()
	if connectErr != nil {
		fmt.Println("Fail to connect database")
	}

	u := User{}
	nameErr := db.Table("user_info").Where("name = ?", username).First(&u).Error
	passwordErr := db.Table("user_info").Where("password = ?", password).First(&u).Error
	if nameErr != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist",
		})
	} else if passwordErr != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 2,
			StatusMsg:  "Wrong password!",
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{0, "Login Success!"},
			UserId:   u.Id,
			Token:    token,
		})
	}

	//if user, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, UserLoginResponse{
	//		Response: Response{StatusCode: 0},
	//		UserId:   user.Id,
	//		Token:    token,
	//	})
	//} else {
	//	c.JSON(http.StatusOK, UserLoginResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	//	})
	//}
}

func UserInfo(c *gin.Context) {

	db, connectErr := common.DBconnection()

	if connectErr != nil {
		println("Fail to connect database!")
	}
	token := c.Query("token")
	user_id := c.Query("user_id")
	u := User{}
	err := common.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			StatusCode: 1,
			StatusMsg:  "Invalid token",
		})
	} else {
		DBerr := db.Table("user_info").Where("id=?", user_id).First(&u).Error
		if DBerr != nil {
			c.JSON(http.StatusInternalServerError, Response{
				StatusCode: 2,
				StatusMsg:  "Unexpected Error!",
			})
		} else {
			c.JSON(http.StatusOK, u)
		}

	}
	//QueryErr := db.Table("user_info").Where("token = ?", token).First(&u).Error
	//if user, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 0},
	//		User:     user,
	//	})
	//} else {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	//	})
	//}
}
