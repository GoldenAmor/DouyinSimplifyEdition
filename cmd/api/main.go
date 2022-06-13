package main

import (
	"dousheng/cmd/api/conn"
	"dousheng/cmd/api/router"
	"dousheng/cmd/api/rpc"
	"github.com/gin-gonic/gin"
)

func main() {
	conn.InitRedis()
	rpc.Init()
	r := gin.Default()
	
	router.InitRouter(r)
	err := r.Run()
	if err != nil {
		return
	}
}
