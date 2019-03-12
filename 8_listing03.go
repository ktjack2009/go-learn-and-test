/*
	log包是多goroutine安全的，这意味着在多个goroutine可以同时调用来自同一个日志记录器的这些函数，
	而不会有彼此间的写冲突
*/

package main

import (
	"log"
)

func init() {
	log.SetPrefix("TRACE: ") // 设置前缀
	/*
		Ldate：日期：2009/01/23
		Ltime：时间：01:23:23
		Lmicroseconds：毫秒级时间：01:23:23.123123【该设置会覆盖Ltime】
		Llongfile：完整路径的文件名和行号：/a/b/c/d.go:23
		Lshortfile：最终文件名元素和行号：d.go:23【覆盖Llongfile】
		LstdFlags：标准日志记录器的初始值
	*/
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	log.Println("message")
	log.Fatalln("fatal message") // 调用完后会接着调用os.Exit(1)
	log.Panicln("panic message") // 调用完后会接着调用panic()
}
