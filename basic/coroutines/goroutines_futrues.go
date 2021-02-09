package coroutines

import (
	"fmt"
	"log"
	"time"
)

type State int

const (
	NEW        State = iota // value --> 0
	RUNNING                 // value --> 1
	COMPLETING              // value --> 2
	CANCELLED               // value --> 3
)

type Task interface {
	Eval(...interface{}) (interface{}, error)
}
type Future interface {
	Get() (interface{}, error)
	GetDuration(time.Duration) (interface{}, error)
	Cancel()
}

func (receiver *FutureTask) Get() (ret interface{}, err error) {
	for {
		select {
		case taskRet := <-receiver.retChan:
			ret = taskRet.ret
			err = taskRet.err
			receiver.retChan = nil
		case state := <-receiver.stateChan:
			log.Printf("get:%d", state)
			if state > RUNNING {
				receiver.stateChan = nil
				receiver.retChan = nil
			}
		}
		if receiver.stateChan == nil && receiver.retChan == nil {
			return ret, err
		}
	}
}
func (receiver *FutureTask) GetDuration(duration time.Duration) (interface{}, error) {
	return nil, nil
}

func (receiver *FutureTask) asyncEval(task Task) {
	go func() {

		evalFunc, err := (task).Eval()
		receiver.retChan <- TaskRet{evalFunc, err}
		receiver.stateChan <- COMPLETING
		fmt.Println("async eval complete")

	}()
}

func (receiver *FutureTask) Cancel() {
	receiver.stateChan <- CANCELLED
	fmt.Println("cancel")

}

type FutureTask struct {
	stateChan chan State
	retChan   chan TaskRet
	state     State
}

type TaskRet struct {
	ret interface{}
	err error
}

func NewFuture(task Task) Future {
	futureTask := FutureTask{
		make(chan State, 1),
		make(chan TaskRet, 1),
		NEW,
	}
	futureTask.asyncEval(task)
	return &futureTask
}
