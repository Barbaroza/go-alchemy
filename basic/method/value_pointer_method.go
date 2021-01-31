package method

import "fmt"

type A struct {
	a int32
}

func (receiver A) valueMethod(aValue int32) {
	receiver.a = aValue
}

func (receiver *A) pointerMethod(aValue int32) {
	receiver.a = aValue
}

func DiffMethodDemo() {
	instance := A{0}
	instancePtr := new(A)

	instance.valueMethod(1)
	fmt.Printf("instance.valueMethod %v \n", instance)
	instance.pointerMethod(2)
	fmt.Printf("instance.pointerMethod %v \n", instance)

	(&instance).pointerMethod(3)
	fmt.Printf("(&instance).pointerMethod %v \n", instance)

	instancePtr.valueMethod(1)
	fmt.Printf("instancePtr.valueMethod %v \n", instancePtr)
	instancePtr.pointerMethod(2)
	fmt.Printf("instancePtr.pointerMethod %v \n", instancePtr)

	(*instancePtr).valueMethod(3)
	fmt.Printf("(&instancePtr).pointerMethod %v \n", instancePtr)

}
