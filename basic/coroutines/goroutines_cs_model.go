package coroutines

import "fmt"

type Response interface {
}
type Request interface {
}

type RequestWrapper struct {
	Request
	process func(request Request) Response
	rspChan chan Response
}

func (receiver *RequestWrapper) processExtender() {
	receiver.rspChan <- receiver.process(receiver.Request)
}

func server() chan RequestWrapper {

	reqChan := make(chan RequestWrapper)
	go func() {
		for {
			requestWrapper := <-reqChan
			go requestWrapper.processExtender()
		}
	}()

	return reqChan
}

func CsModelDemo() {
	wrappers := server()
	for i := 0; i < 100; i++ {
		temp := i
		requestWrapper := RequestWrapper{
			rspChan: make(chan Response),
			Request: temp,
			process: func(req Request) Response {
				fmt.Printf("%d process start \n", temp)
				fmt.Printf("%d process end \n", temp)
				return req
			},
		}
		wrappers <- requestWrapper
	}
}
