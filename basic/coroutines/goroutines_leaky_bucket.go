package coroutines

var clientChan chan int8 = make(chan int8, 100)
var serverChan chan int8 = make(chan int8, 100)

func leakyBucketClient() {
	for {
		var buffer int8
		// Grab a buffer if available; allocate if not
		select {
		case buffer = <-clientChan:
		default:
			// None free, so allocate a new one
			buffer = 0
		}
		serverChan <- buffer
	}
}

func leakyBucketServer() {
	for {
		b := <-serverChan
		// Wait for work.
		// process data
		select {
		case clientChan <- b:
			// Reuse buffer if free slot on freeList; nothing more to do
		default:
			// Free list full, just carry on: the buffer is 'dropped'

		}
	}
}
