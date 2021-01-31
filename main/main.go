package main

import (
	"context"
	"fmt"
	"time"
)

func text(ctx context.Context, exitchan chan int, i int) {
	defer func() { exitchan <- 1 }()
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
	exitchan := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		go text(ctx, exitchan, i)
	}

	time.Sleep(time.Second)
	cancel()
	fmt.Println(ctx.Err())
	// 退出
	for i := 1; i <= 5; i++ {
		<-exitchan
	}
}
