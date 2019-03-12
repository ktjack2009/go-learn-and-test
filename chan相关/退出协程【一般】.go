package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, cannel chan bool) {
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("你好")
		case <-cannel:
			fmt.Println("退出")
			return
		}
	}
}

func main() {
	cannel := make(chan bool)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg, cannel)
	}
	time.Sleep(time.Microsecond * 10)
	close(cannel)
	wg.Wait()
}
