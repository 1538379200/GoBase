package main

import "fmt"

// 条件判断语句和java的一样，但是不需要;去结尾

func main()  {
	num := 10
	for x:=0; x<=num; x++ {
		if x-3 > 2 {
			fmt.Println(x)
		}else {
			fmt.Println("小鱼")
		}
	}
}
