package coroutines

import "fmt"

func push(strChan chan interface{}, str interface{}) {
	strChan <- str
	fmt.Printf("push:%s\n", str)

}

func poll(strChan chan interface{}) interface{} {
	str := <-strChan
	fmt.Printf("poll:%s\n", str)
	return str
}

func ChanDemo(buffer int) {
	ch := make(chan interface{}, buffer)

	go push(ch, "1")
	go push(ch, "2")
	go push(ch, "3")
	go func() {
		for {
			poll(ch)
		}
	}()
	Wait(3)
}
