package main

import (
	"fmt"
	"time"
)

// 这个文件是在 noMe_GoroutinePool 的启发下寻找的一种更适合自己的方法

type Cases struct {
	f func(...string)
}

func NewCases(f func(...string)) *Cases {
	c := Cases{f: f}
	return &c
}

func (c *Cases) ExCase(msg ...string) {
	c.f(msg...)
}

type CasePool struct {
	maxWorker   int
	maxCount    int
	PushChannel chan *Cases
	JobChannel  chan *Cases
}

func NewCasePool(maxWorker, maxCount int) *CasePool {
	cp := CasePool{
		maxWorker:   maxWorker,
		PushChannel: make(chan *Cases),
		JobChannel:  make(chan *Cases),
		maxCount:    maxCount,
	}
	return &cp
}

func (p *CasePool) worker(msg ...string) {
	for i := range p.JobChannel {
		i.ExCase(msg...)
	}
}

func (p *CasePool) Run(msg ...string) {
	for i := 0; i < p.maxWorker; i++ {
		go p.worker(msg...)
	}

	for task := range p.PushChannel {
		p.JobChannel <- task
	}
	close(p.PushChannel)
	close(p.JobChannel)
}

func MainRun(msg ...string) {
	t := NewCases(func(s ...string) {
		fmt.Println(s, time.Now())
	})

	p := NewCasePool(5, 1)

	go func() {
		for i := 0; i < p.maxCount; i++ {
			p.PushChannel <- t
		}
		//p.PushChannel <- t
	}()
	p.Run(msg...)
}

func main() {
	MainRun("好的", "不好的", "测试数据")
}
