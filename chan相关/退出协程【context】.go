package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			fmt.Println("hello")
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Microsecond)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, &wg)
	}
	time.Sleep(10 * time.Microsecond)
	cancel()
	wg.Wait()
}
