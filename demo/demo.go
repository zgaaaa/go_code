package main

import (
	"fmt"
	"time"
)

func text(stopchan chan bool, exitchan chan int, i int) {
	defer func() { exitchan <- 1 }()
	for {
		select {
		case <-stopchan:
			fmt.Printf("协程%v退出,结束监控\n", i)
			return
		default:
			fmt.Printf("协程%v,正在监控\n", i)
			time.Sleep(time.Second * 2)
		}
	}
}

func main() {
	exitchan := make(chan int, 5)
	stopchan := make(chan bool)
	for i := 1; i <= 5; i++ {
		go text(stopchan, exitchan, i)
	}
	time.Sleep(time.Second)
	close(stopchan)
	for i := 1; i <= 5; i++ {
		<-exitchan
	}
}
