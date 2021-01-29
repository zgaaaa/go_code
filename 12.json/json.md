### json

#### 序列化

```go
package main

import (
	"encoding/json"
	"fmt"
)

//结构体序列化 
func strumarshal() []byte {
	// Monster 结构体
	type Monster struct {
		Name string `json:"name"` // tag标签,可以理解为在序列化时对字段重命名
		Age  int    `json:"age"`
		Sex  string `json:"sex"`
	}
	data := Monster{
		Name: "qwe",
		Age:  14,
		Sex:  "nan",
	}
	json, _ := json.Marshal(&data)
	return json
}

// 切片序列化
func slicemarshal() []byte {
	slice := []map[string]interface{}{}
	map1 := map[string]interface{}{
		"name": "sdfw",
		"sex":  "nan",
		"age":  18,
	}
	slice = append(slice, map1)
	map2 := map[string]interface{}{
		"name": "为如果",
		"sex":  "nv",
		"age":  23,
	}
	slice = append(slice, map2)
	data, _ := json.Marshal(slice)
	return data
}

func main() {
	fmt.Printf("%s\n", strumarshal())
	fmt.Printf("%s\n", slicemarshal())
}
//{"name":"qwe","age":14,"sex":"nan"}
//[{"age":18,"name":"sdfw","sex":"nan"},{"age":23,"name":"为如果","sex":"nv"}]
```



#### 反序列化

```go
package main

import (
	"encoding/json"
	"fmt"
)

func unmarshalslice() interface{} {
	str := `[{"age":18,"name":"sdfw","sex":"nan"},{"age":23,"name":"为如果","sex":"nv"}]`
	// 提前定义好接受反序列化数据的变量
	// 变量类型要和序列化前的数据类型保持一致.
	var slice []map[string]interface{} 
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
```

