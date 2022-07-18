package main

import "fmt"

// select的方法与switch大同小异，但是select是随机抽取满足条件的执行，否则阻塞
// switch是顺序判断
func main()  {
	var talk = "优秀"
	var test = 60
	switch test {
		case 90: talk = "优秀"
		case 80: talk = "很好"
		case 60: talk = "及格"
		default: talk = "努力"
	}
	fmt.Println(talk)	// 输出及格
}
