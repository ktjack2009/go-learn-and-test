package main

import (
	"fmt"
	"runtime"
)

func main() {
	ChanDemo1()
}

func ChanDemo1() {
	ch1 := make(chan interface{}, 2)
	ch1 <- 1
	ch1 <- 2
	close(ch1)

	for {
		a, _ := <-ch1
		// if !ok && a == 0 {
		// 	break
		// } else {
		// 	fmt.Println(a)
		// }
		if a == nil {
			break
		}
		fmt.Println(a)
	}
}

func ChanDemo2() {
	// 多核并行计算
	var Cores = runtime.NumCPU() // cpu个数
	runtime.GOMAXPROCS(1)        // 启动goroutine之前设置运行cpu个数

	baseChan := make(chan uint64, 1024)
	resultChan := make([]chan uint64, Cores)
	// 放入需要计算的数
	go func(ch chan uint64) {
		var i uint64
		for i = 0; i <= 10000000; i++ {
			ch <- i
		}
		close(ch)
	}(baseChan)

	Sum := func(ch1 chan uint64, ch2 chan uint64, id int) {
		var temp uint64 = 0
		for {
			value, ok := <-ch1
			if !ok && value == 0 {
				// 通道关闭且取出的数为0
				break
			}
			temp += value
		}
		fmt.Printf("第【%d】任务最终结果是%d\n", id, temp)
		ch2 <- temp
	}
	for i := 0; i < Cores; i++ {
		resultChan[i] = make(chan uint64)
		go Sum(baseChan, resultChan[i], i)
	}

	var sum uint64 = 0
	for _, valueChan := range resultChan {
		value := <-valueChan
		sum += value
	}
	fmt.Println(sum)
}
