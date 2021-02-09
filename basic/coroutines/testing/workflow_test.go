package testing

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestWorkFlow(t *testing.T) {

	arr := make([]int, 60000)

	for i := 0; i < cap(arr); i++ {
		arr[i] = rand.Int() % 256
	}
	before := time.Now().UnixNano() / 1e6

	sort.Ints(arr)
	var less64, less128, less192, less256 int = 0, 0, 0, 0
	for i := 0; i < 100000; i++ {
		for n := 0; n < cap(arr); n++ {
			temp := arr[n]
			if temp < 64 {
				less64++
			} else if temp < 128 {
				less128++
			} else if temp < 192 {
				less192++
			} else {
				less256++
			}
		}

	}
	after := time.Now().UnixNano() / 1e6

	less128 += less64
	less192 += less128
	less256 += less192

	fmt.Println(after - before)
}
