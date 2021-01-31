package method

import "fmt"

type T struct {
	value int32
}

func (receiver *T) TMethod() {
	//dosomething
}

func (receiver *T) setValue2(value int32) {
	receiver.value = value
}
func (receiver T) setValue1(value int32) {
	receiver.value = value
}

func MethodDemo() {
	t1 := T{}
	t1.TMethod()

	t2 := new(T)
	t2.TMethod()

}

func ReceiverDemo() {
	t := T{0}

	t.setValue1(1)

	fmt.Printf("setValue1: %v \n", t)

	t.setValue2(2)
	T.setValue1(t, 1)
	fmt.Printf("setValue2: %v \n", t)

}
