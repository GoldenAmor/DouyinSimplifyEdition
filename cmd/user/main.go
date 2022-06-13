package main

import (
	"dousheng/cmd/user/conn"
	user "dousheng/kitex_gen/user/userservice"
	"dousheng/pkg/conf"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	conn.InitGorm()

	addr, _ := net.ResolveTCPAddr("tcp", conf.UserPort)
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))

	svr := user.NewServer(new(UserServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
