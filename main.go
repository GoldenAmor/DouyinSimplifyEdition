package main

import (
	"github.com/RaymondCode/simple-demo/conn"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	repository.Init()
	conn.InitGorm()
	r := gin.Default()
	router.InitRouter(r)
	//ConnectDB()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
