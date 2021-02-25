package coroutines

type data struct {
}

func serialProcessData(in <-chan *data, out chan<- *data) {
	for data := range in {
		preProcess(data)
		stepA(data)
		stepB(data)
		postProcess(data)
		out <- data
	}
}

func postProcess(d *data) {

}

func stepB(d *data) {

}

func stepA(d *data) {

}

func preProcess(d *data) {

}
func pre(chan *data, chan *data) {

}

func parallelProcessData(in <-chan *data, out chan<- *data) {
	preChan := make(chan *data)
	aChan := make(chan *data)
	bChan := make(chan *data)
	postChan := make(chan *data)

	pre(preChan, aChan)
	a(aChan, bChan)
	b(bChan, postChan)
	post(postChan, out)

}

func post(postChan chan *data, out chan<- *data) {

}

func a(postChan chan *data, out chan<- *data) {

}

func b(postChan chan *data, out chan<- *data) {

}
