package main

import "fmt"

var map1 map[int]int

func text(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	map1[n] = res
}

func main() {
	for i := 1; i <= 200; i++ {
		go text(i)
	}
	for index, v := range map1 {
		fmt.Printf("%v : %v\n", index, v)

	}
}
