package coroutines

import "fmt"

func GoroutinesDemo2() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := make(chan int)
	go func() {
		sum := 0
		for _, v := range arr {
			sum += v
		}
		res <- sum
	}()
	fmt.Printf("sum res : %d \n", <-res)
}
