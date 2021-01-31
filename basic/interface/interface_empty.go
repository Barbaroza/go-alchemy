package _interface

import "fmt"

type Object interface{}

type Vector struct {
	data []Object
}

func InterfaceEmptyDemo() {

	var obj Object = man{}
	typeFunction := func(any interface{}) {
		switch t := any.(type) {
		case man:
			fmt.Printf("%v is man", t)
		default:
			fmt.Printf("%v is unknown", t)
		}
	}
	typeFunction(obj)
}

func InterfaceEmptyCopy() {
	slice := make([]int, 5)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	slice[3] = 4
	slice[4] = 5
	emptySlice := make([]Object, len(slice))
	for i, d := range slice {
		emptySlice[i] = d
	}
}
