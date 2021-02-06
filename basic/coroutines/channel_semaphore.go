package coroutines

import "fmt"

func Semaphore() {
	ch := make(chan interface{})
	go func() {
		//compute
		ch <- new(int)
	}()

	<-ch
}

func Provide(ch chan interface{}, n int) {
	for i := 0; i < n; i++ {
		ch <- i
		fmt.Printf("provide : %d \n", i)
	}
}

func Consume(ch chan interface{}) {
	for v := range ch {
		fmt.Printf("consume : %d \n", v)
	}
}

func ConsumeProvider() {
	ch := make(chan interface{}, 5)
	go Provide(ch, 10)
	go Consume(ch)
	Wait(5)
}
