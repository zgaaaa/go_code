package demo

import (
	"fmt"
	"strconv"
)

// account 结构体
type account struct {
	zhanghao int
	yue      float64
	mima     int
}

// Getaccount 获取结构体
func Getaccount() *account {
	return &account{}
}

// Setzhanghao 更改账号
func (a *account) Setzhanghao(zhanghao string) {
	zhanghao1, _ := strconv.Atoi(zhanghao)
	if len(zhanghao) >= 6 && len(zhanghao) <= 10 {
		a.zhanghao = zhanghao1
	} else {
		fmt.Println("输入有误")
	}
}

// Getzhanghao 获取账号
func (a *account) Getzhanghao() int {
	return a.zhanghao
}
