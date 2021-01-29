package demo

import (
	"testing"
)

func TestSum(t *testing.T) {
	type TestData struct {
		n1 int
		n2 int
	}
	data := []TestData{
		{1, 2},
		{3, 4},
		{9, 12},
		{-9, 3},
	}
	for _, v := range data {
		result := sum(v.n1, v.n2)
		if result == (v.n1 + v.n2) {
			t.Log("测试正确")
		} else {
			t.Fatal("测试错误")
		}
	}
}
