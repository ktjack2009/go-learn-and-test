// goroutine的使用

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1) // 分配一个逻辑处理器给调度器使用
	var wg sync.WaitGroup

	// wg用来等待程序完成
	// 计数加2，表示要等待两个goroutine
	wg.Add(2)

	fmt.Println("start goroutines")

	go func() {
		// 在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
			time.Sleep(time.Second)
		}

	}()

	go func() {
		// 在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
			time.Sleep(time.Second)
		}
	}()

	fmt.Println("waiting to finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}
