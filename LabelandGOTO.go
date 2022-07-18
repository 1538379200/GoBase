package main

import "fmt"

// lable 是一个命名，给循环在break、continue的时候写上可以针对性停止

func loop() {
	var i = 0
mainloop: // 标明一个label
	for {
		for {
			if i == 12 {
				break mainloop // 停止整个label
			} else {
				fmt.Println(i)
				i += 1
			}
		}
	}
	fmt.Println(i + 1)
	goto gotothis // 写这个就会跳转到第四个, 13过后就是16了
	fmt.Println(i + 2)
	fmt.Println(i + 3)
gotothis:
	fmt.Println(i + 4)
}

func main() {
	loop()
}
