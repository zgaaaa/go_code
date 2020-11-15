package main

import "fmt"

func main() {
	slice1 := [][]int{{1,2},{3,4}}
	fmt.Println(slice1)
	slice2 := slice1
	slice2[0][0] = 5
	fmt.Println(slice1)
}
