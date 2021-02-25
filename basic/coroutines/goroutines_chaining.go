package coroutines

import "fmt"

func link(left chan int, right chan int) {
	a := <-right
	left <- a + 1
}
func ChainingDemo() {
	var mostLeft = make(chan int)
	var left, right chan int = nil, mostLeft

	for i := 0; i < 500; i++ {
		left, right = right, make(chan int)
		go link(left, right)
	}
	right <- 2
	res := <-mostLeft
	fmt.Println(res)
}
