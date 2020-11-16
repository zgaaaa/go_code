package main

import "fmt"

// Circle 圆
type Circle struct {
	radis float64 // 半径
}

func (c *Circle) xiugai() float64 {
	(*c).radis = 10
	return c.radis
}

func main() {
	c1 := Circle{5}
	c1.radis = 20
	c1.xiugai()
	fmt.Println(c1.radis)
}
