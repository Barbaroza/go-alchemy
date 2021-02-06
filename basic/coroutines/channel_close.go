package coroutines

import "fmt"

var ch chan interface{} = make(chan interface{})

func ChannelCloseDemo() {
	var closeableCh chan interface{} = make(chan interface{})
	defer close(closeableCh)
}

func ChannelCheck() {
	if v, ok := <-ch; ok {
		fmt.Printf("%d", v)
	} else {

	}
}

func ChannelCloseDemo2() {

	go func() {
		push(ch, "1")
		push(ch, "2")
		push(ch, "3")
		push(ch, "4")
		close(ch)
	}()
	for v := range ch {
		fmt.Printf("%s \n", v)
	}
}
