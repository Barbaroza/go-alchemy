package _interface

import (
	"fmt"
	"reflect"
)

type animal interface {
	eat()
}
type person interface {
	speak()
}

type man struct {
}

func (receiver man) eat() {
	fmt.Printf("%v eat\n", reflect.TypeOf(man{}))
}
func (receiver man) speak() {
	fmt.Printf("%s speak\n", reflect.TypeOf(man{}).Name())
}

func InterfaceDemo() {
	man := man{}
	var animal = animal(man)
	animal.eat()
	var person = person(man)
	person.speak()

}
