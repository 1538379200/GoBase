package main

import "fmt"

// 数组定义需要指定数量, 如果数量少则用0填充
var arry1 = [10]int{1, 2, 3, 4, 5}

// 可以使用...来进行占位，会自动计算数量
var arry2 = [...]int{1, 2, 3, 4}

func main()  {
	fmt.Println(arry1)		// 输出后面为0
	fmt.Println(arry2)
	arry2 = [...]int{1, 2, 4, 5}
	fmt.Println(arry2)		// 输出1245
	// 切片和数组都是可以使用这两种查看元素数量的
	fmt.Println(cap(arry1)) // 查看容量
	fmt.Println(len(arry1)) 	// 查看长度
}
