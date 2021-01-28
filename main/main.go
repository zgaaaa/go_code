package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file1 := "d:/go_work/github.com/zgaaaa/go_code/go.mod"
	file2 := "d:/go_work/github.com/zgaaaa/go_code/text.txt"
	data, err := ioutil.ReadFile(file1)
	if err != nil {
		fmt.Println("file1  ",err)
	}
	fmt.Printf("%s", data)
	err = ioutil.WriteFile(file2,data,0777) // 覆盖写
	if err != nil {
		fmt.Println("写入错误 ",err)
	}
}
