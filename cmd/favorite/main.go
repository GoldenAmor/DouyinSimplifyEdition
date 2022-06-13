package main

import (
	"dousheng/cmd/favorite/conn"
	favorite "dousheng/kitex_gen/favorite/favoriteservice"
	"dousheng/pkg/conf"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	conn.InitGorm()

	addr, _ := net.ResolveTCPAddr("tcp", conf.FavoritePort)
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))
	svr := favorite.NewServer(new(FavoriteServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
