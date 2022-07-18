package main

import "fmt"

// map的映射和python的字典一样，是一个保存键值对的数据类型

//使用map需要对其进行初始化才能使用，有几种方式去创建并初始化，不初始化为nil映射，是不能传值操作的
// 直接在创建结构的时候，后面加上{}即为一个空的map，空map不是nilmap
var mymap = map[string]string{}

//使用make方式创建
var mymap2 = make(map[string]string)

//如果直接创建则需要去进行初始化，可以使用make，可是直接在后面赋值，如上面两个
var needinit map[string]string

// 创建并给定默认值
var defaultmap = map[string]string{"测试": "好样的"}

func main() {
	mymap["测试"] = "测试一"
	fmt.Println(mymap) // 输出map[测试:测试一]
	mymap2["测试"] = "测试二"
	fmt.Println(mymap2) // 输出 map[测试:测试二]
	needinit = make(map[string]string)
	needinit["测试"] = "测试三"
	fmt.Println(needinit) // 输出 map[测试:测试三]
	fmt.Printf("%T", needinit)
	fmt.Println(defaultmap) // 输出 map[测试:好样的]

	//读取值则和字典的形式一样
	fmt.Println(needinit["测试"]) // 输出 测试三
	// 可以获取map的长度
	fmt.Println(len(mymap))
}
