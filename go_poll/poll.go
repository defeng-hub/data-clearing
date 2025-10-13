package go_poll

import (
	"sync"
)

// Task 定义任务类型（你也可以换成 func(interface{})）
type Task func()

type Pool struct {
	taskChan chan Task
	wg       sync.WaitGroup
}

// NewPool 创建协程池，workers 是最大并发协程数
func NewPool(workers int) *Pool {
	p := &Pool{
		taskChan: make(chan Task),
	}

	for i := 0; i < workers; i++ {
		go p.worker(i)
	}

	return p
}

// Submit 提交任务
func (p *Pool) Submit(task Task) {
	p.wg.Add(1)
	p.taskChan <- task
}

// worker 执行任务
func (p *Pool) worker(id int) {
	for task := range p.taskChan {
		task()
		p.wg.Done()
	}
}

// Wait 等待所有任务完成
func (p *Pool) Wait() {
	p.wg.Wait()
}

// Close 关闭任务通道（通常在 Wait 后调用）
func (p *Pool) Close() {
	close(p.taskChan)
}
