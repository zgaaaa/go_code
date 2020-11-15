package main

import (
	"fmt"
)

// Person 结构体Person
type Person struct {
	Name string
}

// 给结构体Person绑定一个方法
func (p Person) name() {
	fmt.Println(p.Name, "是一个好人")
}

func (p Person) jisuan(n int) {
	sum := 0
	for i := 1; i < n+1; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func (p Person) getsum(a, b int) {
	fmt.Println(a + b)
}

func main() {
	var p1 Person
	p1.Name = "tom"
	p1.jisuan(1000)
	p1.getsum(10, 20)
}
