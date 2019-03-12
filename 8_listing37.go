package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// 创建一个buffer，并将一串字符串写入buffer
	// 使用实现io.Writer的Write方法
	var b bytes.Buffer
	b.Write([]byte("Hello "))

	// 使用Fprintf将一个字符串拼接到buffer里
	fmt.Fprintf(&b, "World")
	b.WriteTo(os.Stdout)
}
