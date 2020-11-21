package main

import (
	"fmt"
	"strings"
	"time"
)

func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else {
			start += index + 1
			end += index + 1
		}
	}
	return end - start
}

func main() {
	startime := time.Now()
	str1 := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(str1))
	endtime := time.Since(startime)
	fmt.Println(endtime)
}
