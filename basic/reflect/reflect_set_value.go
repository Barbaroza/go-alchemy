package reflect

import (
	"fmt"
	"reflect"
)

func ReflectSetDemo() {
	var index int = 2
	fmt.Printf("before reflect set ,index value:%d\n", index)

	reflectValue := reflect.ValueOf(&index)
	ok := reflectValue.CanSet()
	fmt.Printf("reflect.ValueOf(&index).CanSet():%v\n", ok)

	if ok := reflectValue.Elem().CanSet(); ok {
		reflectValue = reflectValue.Elem()
		reflectValue.SetInt(1231)
	}
	fmt.Printf("reflect.ValueOf(&index).Elem().CanSet():%v\n", ok)


	fmt.Printf("after reflect set ,index value:%d\n", index)

}
