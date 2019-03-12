package main

import (
	"log"
	"net/rpc"
)

type ArithRequest struct {
	A int
	B int
}

type ArithResponse struct {
	Pro int // 乘积
	Quo int // 商
	Rem int // 余数
}

func main() {
	conn, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalln("dailing error:", err)
	}
	req := ArithRequest{9, 2}
	var reply int
	err = conn.Call("Arith.Multiply", req)
}
