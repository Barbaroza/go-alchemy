package coroutines

import "fmt"

// Send the sequence 2, 3, 4, ... to returned channel
func generate2() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
			fmt.Printf("generate2 push %d \n", i)

		}
	}()
	return ch
}

// @star  Filter out input values divisible by 'prime', send rest to returned channel
func filter2(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				fmt.Printf("filter2  poll %d \n", i)

				out <- i
			}
		}
	}()
	return out
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate2()
		for {
			prime := <-ch
			fmt.Printf("before filter2  poll %d \n", prime)
			ch = filter2(ch, prime)
			out <- prime
		}
	}()
	return out
}

func PrimeDemo2() {
	primes := sieve()
	for {
		fmt.Printf("res : %d\n", <-primes)
	}
}
