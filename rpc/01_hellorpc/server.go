package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

// Go语言的RPC规则：方法只能有两个可序列化的参数，其中第二个参数是指针类型
// 并且返回一个erro类型，同时必须是公开的方法
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	// 注册为一个rpc服务
	rpc.RegisterName("HelloService", new(HelloService))
	listener, _ := net.Listen("tcp", ":1234")
	for {
		conn, _ := listener.Accept()
		rpc.ServeConn(conn)
	}
}
