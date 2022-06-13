package main

import (
	"dousheng/cmd/video/conn"
	video "dousheng/kitex_gen/video/videoservice"
	"dousheng/pkg/conf"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	conn.InitGorm()

	addr, _ := net.ResolveTCPAddr("tcp", conf.VideoPort)
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))

	svr := video.NewServer(new(VideoServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
