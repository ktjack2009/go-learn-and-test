// 使用recover()函数来捕捉异常
package main

import (
	"fmt"
	"log"
)

func main() {
	foo()
	foo1()
}

func foo() {
	/*
		无论foo()中是否触发了错误处理流程，该匿名defer函数都将在函数退出时得到执行。
		假如foo()中触发了错误处理流程，recover()函数执行将使得该错误处理过程终止
	*/
	defer func() {
		if err := recover(); err != nil {
			log.Println("出现异常")
			log.Println(err)
		}
	}()

	var a, b int
	a = 1
	b = 0
	fmt.Println(a / b)
}

func foo1() {
	fmt.Println("demo2运行")
}
