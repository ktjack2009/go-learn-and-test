// 展示使用无缓冲通道模拟4个goroutine直接的接力赛
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// 创建无缓冲通道
	baton := make(chan int)
	wg.Add(1)

	go Runner(baton)

	baton <- 1
	wg.Wait()

}

func Runner(baton chan int) {
	var newRunner int
	runner := <-baton
	fmt.Printf("Runner %d Running\n", runner)

	// 创建下一位跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d to the line\n", newRunner)
		go Runner(baton)
	}

	// 围绕跑道跑
	time.Sleep(1000 * time.Millisecond)

	// 比赛结束了吗
	if runner == 4 {
		fmt.Printf("Runner %d Finished\n", runner)
		wg.Done()
		return
	}

	// 将接力棒交给下一位跑步者
	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)
	baton <- newRunner
}
