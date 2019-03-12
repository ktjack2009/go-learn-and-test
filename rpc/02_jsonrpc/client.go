package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:1234")
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	var reply string
	client.Call("HelloService.Hello", "hello", &reply)
	fmt.Println(reply)
}
