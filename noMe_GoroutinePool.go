package main

import (
	"fmt"
	"time"
)

// 定义一个结构体用来存放运行的函数

type Task struct {
	f func() error
}

// 建立一个构造函数初始化结构体

func NewTask(f func() error) *Task {
	t := Task{
		f: f,
	}
	return &t
}

// 定义一个结构体的方法，运行这个函数

func (t *Task) Execute() {
	t.f()
}

// 定义三个数据
// maxWork用于确定最多开多少个
// push的通道用于将外部任务推送至Job中
// Job通道用于将任务传输出来进行运行操作

type Pool struct {
	maxWorker   int
	PushChannel chan *Task
	JobChannel  chan *Task
}

// 创建一个构造函数初始化协程池结构体

func NewPool(maxWorker int) *Pool {
	p := Pool{
		maxWorker:   maxWorker,
		PushChannel: make(chan *Task),
		JobChannel:  make(chan *Task),
	}
	return &p
}

// 整一个推送任务进Job并进行运行的函数

func (p *Pool) StartWorker(workerID int) {
	// 不断从JobChannel中取出储存的函数，进行运行
	for task := range p.JobChannel {
		task.Execute()
		fmt.Println("任务: ", workerID, "执行任务完毕")
	}
}

// 写一个协程池的运行函数，即推送数据

func (p *Pool) Run() {
	for i := 0; i < p.maxWorker; i++ {
		go p.StartWorker(i)
	}
	// 从Push管道，将数据推送到Job中去
	for task := range p.PushChannel {
		p.JobChannel <- task
	}

	close(p.JobChannel)
	close(p.PushChannel)
}

func main() {
	t := NewTask(func() error {
		fmt.Println(time.Now().Format("2006/01/02 15:04:05"))
		return nil
	})
	p := NewPool(3)
	go func() {
		for i := 0; i < 10; i++ {
			p.PushChannel <- t
		}
	}()

	p.Run()
}
