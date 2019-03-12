package main

import (
	"errors"
	"net"
	"net/rpc"
)

type Arith struct {
}

type ArithRequest struct {
	A int
	B int
}

type ArithResponse struct {
	Pro int // 乘积
	Quo int // 商
	Rem int // 余数
}

func (_ *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

func (_ *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("divide by zero")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":1234")
	listener, _ := net.ListenTCP("tcp", tcpAddr)
	for {
		conn, _ := listener.Accept()
		rpc.ServeConn(conn)
	}
}
