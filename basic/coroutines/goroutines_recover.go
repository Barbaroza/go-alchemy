package coroutines

import (
	"log"
	"time"
)

func Panic() {
	panic("error occurred")
}

func working(w interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("%s \n", err)
		}
	}()
	log.Printf("working: %s \n", w)
	Panic()
}

func GoroutinesRecoverDemo() {

	queue := make(chan interface{}, 10)

	go func() {
		for i := 0; i < 20; i++ {
			queue <- i
		}
	}()
	go func() {
		for w := range queue {
			working(w)
		}
	}()

	time.Sleep(10e9)
}
