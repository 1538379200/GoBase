package main

import (
	"fmt"
)

// 第一个括号为形参，第二个括号为返回值
// 如果返回值写了一个变量表示，则在下方可以直接使用，并且return可以省略返回的变量名
func test1(name string) (desc string) {
	desc = "我的名字叫：" + name
	return
}

// 可以多个参数写后面写一个类型，表示前面都是这个类型
// 返回值没有写变量的，return的返回值必须写
// 上面没有写desc变量，所以下面需要创建变量，使用 :=
func test2(name, gender string, age int) string {
	desc := fmt.Sprintf("我的名字是：%v， 我的性别是：%v，我的年龄是：%d", name, gender, age)
	return desc
}

// 不定长参数，用...表示，可以传入0个或者1个、多个，不定长的参数必须放到最后，多个参数，是用一个切片包裹的
func test3(name string, hobby ...string) {
	desc := fmt.Sprintf("我的名字是：%v，我的爱好有：%v", name, hobby)
	fmt.Println(desc)
}

// 下面演示构造函数和方法
// go并不是完全面向对象的编程语言，所以没有类的概念，但是可以用struct模拟类的概念，专门给一个类型使用的函数称之为方法

// 多个方法即可组成一个接口
type testinternal interface {
	useperson() string
}

// 定义一个结构体
type person struct {
	name   string
	gender string
	age    int
}

// 使用一个构造函数，初始化结构体数据，返回一个person的指针值
func NewPerson(name, gender string, age int) *person {
	return &person{
		name:   name,
		gender: gender,
		age:    age,
	}
}

// 方法名前面加上括号，写上的是针对那个类型使用的方法，这里是person
func (p person) useperson() string {
	desc := fmt.Sprintf("我的名字是：%v，我的性别是：%v，我的年龄是：%v", p.name, p.gender, p.age)
	return desc
}

func main() {
	fmt.Println(test1("张三"))
	fmt.Println(test2("张三", "男", 19))
	test3("张三", "吃饭", "睡觉", "打麻将")

	// 使用useperson方法
	// New实例一下结构体，初始化数据
	personel := NewPerson("李四", "女", 30)
	// 在这个初始化的struct下面有useperson方法，就可以直接使用，相当于python类下面的def方法
	fmt.Println(personel.useperson())
	// 通过接口调用
	var personel2 testinternal = person{"王五", "男", 29}
	fmt.Println(personel2.useperson())
}
