package coroutines

import (
	"fmt"
	"time"
)

var duration time.Duration = 1e8
var duration2 time.Duration = 1e9

var ticker = time.Tick(duration)
var after = time.After(duration2)

var counts = 10

func ChannelTickerDemo() {
	temp := counts
	for counts > 0 {
		counts--
		t := <-ticker
		fmt.Printf("%d : %v\n", temp-counts, t)
	}
}

func ChannelAfterDemo() {
	for {
		select {
		case <-ticker:
			fmt.Println("ticker")

		case <-after:
			fmt.Println("after")
			fmt.Println("demo exit")
			return
		}
	}
}
