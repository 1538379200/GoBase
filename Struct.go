package main

import "fmt"

// 定义一个结构体
type mystruct struct {
	name string
	age  int
}

func main() {
	var struct1 mystruct // 定义一个变量struct1，类型指定为mystruct的样式
	struct1.name = "测试"
	struct1.age = 18
	fmt.Println(struct1)      // 输出 {测试 18}
	fmt.Println(struct1.name) // 使用.的方式可以获取里面的值，和python字典取值有点像  输出 测试
	// 下面为匿名结构体
	struct1 = struct {
		name string
		age  int
	}{name: "好的", age: 90}
	fmt.Println(struct1) // 直接使用这种方式全部赋值，其实相当于重新写了一个struct
	// 也可以使用下面这种直接创建变量并传入数据，如果一个值不填写，会使用默认的值，int即为0
	strut2 := mystruct{name: "你好"}
	fmt.Println(strut2)
}
