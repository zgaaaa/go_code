package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1,2,3,4,5}
	fmt.Println(runningSum(slice1))
}
func runningSum(nums []int) []int {
    count := 0
    var nums2 []int
    nums2 = make([]int,len(nums))
    for i := 0;i<len(nums);i++ {
        count += nums[i]
        nums2[i] = count
        
    }
    return nums2
}