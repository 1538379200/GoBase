package main

import "fmt"

// 在go中，可以都看作是一个结构，之后再使用这个结构进行赋值等
//这里是定义一个空的切片
var slice1 []int

// 定义初始值
var slice2 = []int{1, 2, 3, 4}

// 定义一个数组，使用切片获取
var arry = [...]int{2, 4, 6, 8}

// 使用make方式创建，可以给定一个长度，是数组的长度，也是切片的初始长度
var arrymake = make([]int, 10)

// 定义一个长度为0的切片，那么这里就是一个空的切片
var arry0 = make([]int, 0)


func main() {
	slice1 = append(slice1, 1, 2, 3, 4, 5)
	fmt.Println(slice1)		// 使用append添加元素，输出
	fmt.Println("slice2的输出是：", slice2)
	slice3 := arry[0:3]
	fmt.Println("slice3的输出是：", slice3)		// 输出2 4 6
	slice3 = append(slice3, 1, 3, 5, 7)
	fmt.Println("进行append后的slice3是：", slice3)
	fmt.Println(arrymake)
	arrymake2 := append(arrymake, 1, 2, 3, 4)
	fmt.Println(arrymake2)		// 输出 [0 0 0 0 0 0 0 0 0 0 1 2 3 4]
	fmt.Println("arry0的数据是：", arry0)
	arry0 = append(arry0, 1, 3, 4, 5)
	fmt.Println(arry0)		// 即使为0，也可以往里面加入数据
}
