/*
	1. goroutine由Go运行时(runtime)管理

*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	defer func() {
		fmt.Println("主程序退出")
	}()

	//GoroutineDemo1()
	//GoroutineDemo2()
	//GoroutineDemo3()
	//GoroutineDemo4()
	GoroutineDemo5()
}

func GoroutineDemo1() {
	// 使用sync.WaitGroup来等待goroutine全部完成
	defer func() {
		fmt.Println("GoroutineDemo1结束")
	}()

	var wg sync.WaitGroup
	wg.Add(10)

	add := func(n int) {
		defer wg.Done()
		t := rand.Intn(10)
		fmt.Printf("协程%d执行【需要%d秒】...\n", n, t)
		time.Sleep(time.Duration(t) * time.Second)
		fmt.Printf("协程%d完成...\n", n)
	}

	for i := 0; i < 10; i++ {
		go add(i)
	}
	wg.Wait()
}

func GoroutineDemo2() {
	// 使用管道来等待goroutine
	defer func() {
		fmt.Println("GoroutineDemo2结束")
	}()

	ch := make(chan int)

	Count := func(ch chan int, id int) {
		fmt.Printf("Counting任务【%d】进行中\n", id)
		ch <- id
	}
	for i := 0; i < 10; i++ {
		go Count(ch, i)
	}

	for i := range ch { // 如果没有关闭通道，会始终尝试从ch中读取数据，导致死锁
		fmt.Printf("[%d]，gogogo", i)
	}
}

func GoroutineDemo3() {
	// 使用管道来等待goroutine
	defer func() {
		fmt.Println("GoroutineDemo3结束")
	}()

	// 设置超时管道
	chTime := make(chan bool, 1)
	go func() {
		time.Sleep(1 * time.Second)
		chTime <- true
	}()

	ch := make(chan int)
	Count := func(ch chan int, id int) {
		fmt.Printf("Counting任务【%d】进行中\n", id)
		ch <- id
	}
	for i := 0; i < 10; i++ {
		go Count(ch, i)
	}

loop:
	for {
		select {
		case i := <-ch:
			fmt.Printf("【%d】完成\n", i)
		case <-chTime: // 从超时管道中取到了数据就跳出循环
			break loop
		}
	}
}

func GoroutineDemo4() {
	// 单独使用一个[]chan int来表示任务
	// 每个元素都是不带缓冲的chan，会阻塞主程序对于管道的读取

	defer func() {
		fmt.Println("GoroutineDemo4结束")
	}()
	Count := func(ch chan bool, id int) {
		fmt.Printf("Counting任务【%d】进行中\n", id)
		ch <- true
	}

	chs := make([]chan bool, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan bool)
		go Count(chs[i], i)
	}

	for _, ch := range chs {
		<-ch
	}
}

func GoroutineDemo5() {
	c1 := make(chan struct{})
	a := func() {
		fmt.Println("c1")
		c1 <- struct{}{}
	}
	go a()
	go a()
	go a()

	for i := 0; i < 3; i++ {
		<-c1
	}
}
