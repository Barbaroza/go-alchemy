package _interface

import "fmt"

type machine interface {
	silence()
}

func AssertType() {
	p := person(new(man))

	if t, ok := p.(machine); ok {
		fmt.Println("The type of areaIntf is: %T\n", t)
	}
	Classifier(p)
}
