package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  // 要使用的goroutine数量
	taskLoad         = 10 // 要处理的工作数量
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// 创建一个带缓冲的通道来管理工作
	tasks := make(chan string, taskLoad)

	// 启动goroutine来处理工作
	wg.Add(numberGoroutines)

	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// 增加一组要完成的工作
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	// 当所有工作处理完成时关闭通道，以便所有goroutine退出
	// 关闭通道仍然能取数据，但是不能向通道发送数据
	close(tasks)
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {	// 检查通道是否关闭且没有数据
			fmt.Printf("Worker: %d shutting Done\n", worker)
			return
		}
		// 显示要开始工作了
		fmt.Printf("Worker: %d started\n", worker)

		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * 10 * time.Millisecond)

		fmt.Printf("Worker: %d Completed %s\n", worker, task)
	}
}
