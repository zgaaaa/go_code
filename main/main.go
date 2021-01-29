package main

import (
	"encoding/json"
	"fmt"
)

func unmarshalslice() interface{} {
	str := `[{"age":18,"name":"sdfw","sex":"nan"},{"age":23,"name":"为如果","sex":"nv"}]`
	var slice []map[string]interface{} // 提前定义好接受反序列化数据的变量
	json.Unmarshal([]byte(str), &slice) // 必须传入指针
	return slice
}

func unmarshalstruct() interface{} {
	str := `{"name":"qwe","age":14,"sex":"nan"}`
	type Monster struct {
		Name string `json:"name"` // tag标签,可以理解为在序列化时对字段重命名
		Age  int    `json:"age"`
		Sex  string `json:"sex"`
	}
	var monster Monster
	json.Unmarshal([]byte(str), &monster) // 必须传入指针
	return monster
}

func main() {
	fmt.Println(unmarshalslice())
	fmt.Println(unmarshalstruct())

}
