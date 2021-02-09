package testing

import (
	"basic/coroutines"
	"fmt"
	"testing"
	"time"
)

type MyTask struct {
}

func (receiver *MyTask) Eval(...interface{}) (interface{}, error) {
	fmt.Println("myTask start ")
	time.Sleep(4e9)
	fmt.Println("myTask end ")
	return 0, nil
}
func TestFuture(t *testing.T) {
	m := new(MyTask)
	var future = coroutines.NewFuture(m)
	time.Sleep(2e9)

	get, _ := future.Get()

	fmt.Printf("future get: %v  \n", get)

}
