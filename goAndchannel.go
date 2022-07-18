package main

import "fmt"

/* GO最得意的就是并发协程了，可以开任意多个，我们使用的时候需要使用管道，
对于一些对于内存的操作，不使用管道将出问题死锁等
如果不使用管道，我们就只能使用sync来添加锁，但是在go中不建议这么做*/

func goroute(c chan int) {
	for i := 0; i < 100001; i++ {
		c <- i // 这个表示往管道添加数据，把i添加到管道c，注意只有这种箭头写法
	}
	close(c) // 在使用完管道后必须关闭管道
}

func main() {
	var cc = make(chan int) //var cc chan int= make(chan int) // 定义一个同类型的管道，作为上述的参数c，记得用make初始化，不然是一个nil的通道，会报错
	go goroute(cc)          // 使用协程方式运行，channel会自动处理内存使用和阻塞问题
	//for {
	//	v, ok := <-cc // 使用两个参数接收，接收完成后，会把OK变成true
	//	if ok {
	//		fmt.Println(v)
	//	} else {
	//		break
	//	}
	//}
	// 下面使用内置的range方式，循环cc管道的所有数据，更简洁
	for v := range cc {
		fmt.Println(v)
	}

}
