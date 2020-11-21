package main

import (
	"fmt"
	"go_code/Go语言基础/demo"
)

func text() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err = ", err)
		}
	}()
	stu1 := demo.Getstudent("tom", 21)
	fmt.Println(*stu1)
	//fmt.Println(stu1.age)
	fmt.Println(stu1.Getage())
}

func main() {
	text()
}
