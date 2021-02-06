package coroutines

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

type task struct {
	fun func() interface{}
}

type worker struct {
	name  string
	queue chan task
	done  chan interface{}
}

type pool struct {
	queue   chan task
	done    chan interface{}
	workers []*worker
}

func (receiver *pool) submit(fun func() interface{}) {
	t := task{fun: func() interface{} {
		return fun()
	}}

	receiver.queue <- t

}

func (receiver *worker) work() {
	go func() {
		for {
			select {
			case task := <-receiver.queue:
				fmt.Printf("%s start\n", receiver.name)
				receiver.done <- task.fun()
				fmt.Printf("%s end\n", receiver.name)
			}
		}
	}()
}
func NewPool(poolSize int) pool {
	var queue = make(chan task, math.MaxInt8)
	var done = make(chan interface{}, math.MaxInt8)
	var workers = make([]*worker, poolSize)
	for i := 0; i < poolSize; i++ {
		worker := worker{queue: queue, done: done, name: "work" + strconv.Itoa(i)}
		workers[i] = &worker
		worker.work()
	}
	return pool{queue: queue, workers: workers, done: done}
}

func WorkerDemo() {
	newPool := NewPool(2)
	for i := 0; i < 10; i++ {
		temp := i
		newPool.submit(func() interface{} {
			fmt.Printf("job%d start\n", temp)
			time.Sleep(1e9)
			fmt.Printf("job%d end\n", temp)
			return temp
		})
	}
	time.Sleep(15e9)
}
