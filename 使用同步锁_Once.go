/*
	两种类型锁：
	1. sync.Mutex：暴力锁，同一时刻只有一个goroutine可以访问
	2. sync.RWMutex：单写多读锁
*/

package main

import (
	"fmt"
	"sync"
)

var x = 0
var wg sync.WaitGroup
var l sync.Mutex
var Once sync.Once

func main() {
	var n = 1000
	for i := 0; i < n; i++ {
		wg.Add(1)
		go MutexDemo()
	}
	wg.Wait()
	fmt.Printf("最终x的值=%d", x)
}

func MutexDemo() {
	l.Lock()
	defer l.Unlock()

	Once.Do(func() {
		fmt.Println("全局只执行一次")
	})
	x = x + 1
	wg.Done()
}
