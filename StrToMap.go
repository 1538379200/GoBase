package main

import (
	"encoding/json"
	"fmt"
)

// 字符串和map形式的互相转换，使用json模块的反序列化操作

func strMap(str string) {
	a := make(map[string]string)
	//json.Marshal()	序列化，转换成json
	erro := json.Unmarshal([]byte(str), &a)
	if erro != nil {
		println("运行错误")
	}
	for k, v := range a {
		fmt.Printf("%v, %v\n", k, v)
	}
}

func main() {
	// ``反引号的作用是不进行转换，输出原数据，这里用来指示一个字符串
	strMap(`{"a": "123"}`)
}
