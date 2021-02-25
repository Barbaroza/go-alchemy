package coroutines

import (
	"fmt"
	"runtime"
	"time"
)

var NCPU int = 4

func MultiCoresDemo() {
	runtime.GOMAXPROCS(NCPU)
	fmt.Printf("numCPU : %d\n", runtime.NumCPU())

	sed := make(chan int, NCPU/2)
	for i := 0; i < NCPU*2; i++ {
		go process(sed)
	}
	time.Sleep(3e9)
	fmt.Println(time.Now().Nanosecond())
}

func process(sed chan int) {
	sed <- 1

	//process
	time.Sleep(1e9)

	fmt.Println(time.Now().Nanosecond())
	<-sed
}
