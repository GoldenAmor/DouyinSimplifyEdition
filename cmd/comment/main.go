package main

import (
	"dousheng/cmd/comment/conn"
	comment "dousheng/kitex_gen/comment/commentservice"
	"dousheng/pkg/conf"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	conn.InitGorm()

	addr, _ := net.ResolveTCPAddr("tcp", conf.CommentPort)
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))

	svr := comment.NewServer(new(CommentServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
