package main

import (
	"context"
	"fmt"
	"time"
)

func text(ctx context.Context, exitChan chan int, i int) {
	defer func() { exitChan <- 1 }()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("协程%v退出,结束监控\n", i)
			return
		default:
			fmt.Printf("协程%v,正在监控\n", i)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second/2)
	exitChan := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		go text(ctx, exitChan, i)
	}
	time.Sleep(time.Second)
	cancel()
	fmt.Println(ctx.Err())
	// 退出
	for i := 1; i <= 5; i++ {
		<-exitChan
	}
}
