package reflect

import (
	"fmt"
	"reflect"
)

type anStructure struct {
	a int32
	b float32
	c string
}

func (receiver anStructure) PrintFields() {
	fmt.Printf("a:%d b:%f c:%s\n", receiver.a, receiver.b, receiver.c)
}

func ReflectStructureDemo() {
	structure := anStructure{1, 3.2, "hey"}
	valueOfStruct := reflect.ValueOf(structure)
	typeOfStruct := reflect.TypeOf(structure)
	fmt.Printf("typeOfStruct:%v\n", typeOfStruct)
	fmt.Printf("valueOfStruct:%v\n", valueOfStruct.Kind())
	for i := 0; i < valueOfStruct.NumField(); i++ {
		fmt.Printf("%v \n", valueOfStruct.Field(i))
	}

	valueMethod := valueOfStruct.Method(0)


	valueMethod.Call(nil)

}
