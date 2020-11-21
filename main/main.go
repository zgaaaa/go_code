package main

import "fmt"

// Calcuator 矩形
type Calcuator struct {
	arr [][]int
}

func (cal Calcuator) jiujiu(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}

func (cal Calcuator) zuanzhi() {
	for i := 0; i < len(cal.arr); i++ {
		for j := 0; j < len(cal.arr[0]); j++ {
			fmt.Print(cal.arr[j][i])
		}
		fmt.Println()
	}
}

func main() {
	var Cal Calcuator
	//Cal.arr = make([][]int, 0)
	Cal.arr = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//Cal.jiujiu(5)
	Cal.zuanzhi()
}
