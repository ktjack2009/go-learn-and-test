package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func main() {
	wg.Add(2)
	go incCounter3(1)
	go incCounter3(2)
	wg.Wait()
	fmt.Println("finally counter:", counter)
}

func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		value := counter

		// 当前goroutine从线程退出，并放回到队列
		// 如果不退出，goroutine仍然会顺序执行
		runtime.Gosched()

		value ++
		counter = value
	}
}

func incCounter2(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		// 绝对安全的加1
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}

func incCounter3(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		// 同一时刻只允许一个goroutine进入这个临界区
		mutex.Lock()
		{
			value := counter
			value ++
			counter = value
		}
		mutex.Unlock()
	}

}
