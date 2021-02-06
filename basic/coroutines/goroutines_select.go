package coroutines

import (
	"fmt"
	"time"
)

var ch1 chan interface{} = make(chan interface{})
var ch2 chan interface{} = make(chan interface{})

func GoroutinesSelectDemo() {
	go func() {
		push(ch1, "1")
		push(ch1, "2")
		push(ch1, "3")
		push(ch1, "4")
		push(ch1, "5")
	}()
	go func() {
		push(ch2, "a")
		push(ch2, "b")
		push(ch2, "c")
		push(ch2, "d")
		push(ch2, "e")
	}()

	go func() {
		for {
			select {
			case v := <-ch1:
				fmt.Printf("poll from ch1 %s \n", v)
			case v := <-ch2:
				fmt.Printf("poll from ch2 %s \n", v)

			}
		}
	}()
	time.Sleep(1e9)
}
