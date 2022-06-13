package main

import (
	"dousheng/cmd/relation/conn"
	relation "dousheng/kitex_gen/relation/relationservice"
	"dousheng/pkg/conf"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	conn.InitGorm()

	addr, _ := net.ResolveTCPAddr("tcp", conf.RelationPort)
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))

	svr := relation.NewServer(new(RelationServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
