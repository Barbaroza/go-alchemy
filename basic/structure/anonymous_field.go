package structure

import (
	"fmt"
	"reflect"
)

type outer struct {
	int32 "anonymous"
	a     int32
	inner "anonymous"
}
type inner struct {
	b int32
}

//Go 语言中的继承是通过内嵌或组合来实现的。

func ExtendsDemo() {
	in := new(outer)
	in.b = 1
	in.a = 2
	typeInfo := reflect.TypeOf(*in)

	count := typeInfo.NumField()
	for i := 0; i < count; i++ {
		fmt.Printf("outer field%d name: %v\n", i, typeInfo.Field(i).Name)
	}

	in2 := outer{10, 10, inner{b: 20}}

	fmt.Printf("new outer:%v", in2)
}
