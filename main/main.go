package main

import (
	"fmt"
	"go_code/Go语言基础/demo"
)

func main() {
	acc := demo.Getaccount()
	acc.Setzhanghao("123456")
	fmt.Println((*acc).Getzhanghao())
}
