package main

import "fmt"

var num int = 10
// 定义一个int类型的指针
var p *int

// 使用new方法创建一个指针
var str = new(string)

func main()  {
	// 使用&可以获取变量的指针地址
	p = &num
	fmt.Println(p)
	// 使用*可以取出指针地址的值
	fmt.Println(*p)
	*p = 80
	fmt.Println(p)
	fmt.Println(*p)
	num = 20
	fmt.Println(p)
	fmt.Println(*p)

	*str = "测试new创建"
	fmt.Println(str)
	fmt.Println(*str)

}
