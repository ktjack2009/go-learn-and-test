package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, _ := rpc.Dial("tcp", "localhost:1234")
	var reply string
	client.Call("HelloService.Hello", "world", &reply)
	fmt.Println(reply)
}
