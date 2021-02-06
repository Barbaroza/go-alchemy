package coroutines

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

func Wait(seconds int64) {
	fmt.Printf("beginning Wait for %d seconds\n", seconds)
	time.Sleep(time.Duration(seconds * 1e9))
	fmt.Printf("end Wait for %d seconds\n", seconds)
}

func GoroutinesDemo(processors int) {
	fmt.Printf("GoroutinesDemo(%d) start\n", processors)
	var numCores = flag.Int("n", processors, "number of CPU cores to use\n")
	runtime.GOMAXPROCS(*numCores)
	go Wait(5)
	go Wait(9)
	Wait(10)
	fmt.Printf("GoroutinesDemo(%d) end\n", processors)

}
