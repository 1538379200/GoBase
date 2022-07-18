package main

import "fmt"

var myarry = [...]int{1, 2, 3, 4, 5, 6, 7, 8}

func main()  {
	// 在数组、切片、string等循环会输出索引和值
	for key, value := range myarry{
		fmt.Println(key, value)
	}

	// for循环可以给定条件进行循环，从0输出到10
	for x:=0; x<=10; x++{
		fmt.Println(x)
	}

	// 如果后面什么条件都没有，相当于无限循环
	num := 1
	for {
		num ++
		if num > 10 {
			break
		}
	}
	fmt.Println(num)

	// 循环的key和value是可以省略的，go中省略和python一样，使用_，这里演示只取值
	for _, value := range myarry{
		fmt.Println(value)
	}
}
