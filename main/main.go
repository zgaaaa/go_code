package main

import "fmt"

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err = ", err)
		}
	}()
	type Cat struct {
		array1 [3]int
		map1   map[string]string
		slice1 []string
	}
	var cat1 Cat
	// cat1.array1 = [3]int{1, 2, 3}
	// cat1.slice1 = []string{"大毛"}
	// cat1.map1 = map[string]string{
	// 	"name":"二毛",
	// }
	cat1.map1["age"] = "4"
	fmt.Println(cat1)
}
