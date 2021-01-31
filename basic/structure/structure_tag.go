package structure

import (
	"fmt"
	"reflect"
)

type Goods struct {
	color   int32 "颜色"
	quality int32 "质量"
}

func ReflectTagInfo() {
	goods := Goods{color: 1, quality: 20}
	field1 := reflect.TypeOf(goods).Field(0)
	field2 := reflect.TypeOf(goods).Field(1)
	fmt.Printf("goods:%v \n", goods)
	fmt.Printf("field1 name :%v tag: %v \n", field1.Name, field1.Tag)
	fmt.Printf("field2 name :%v tag: %v \n", field2.Name, field2.Tag)
}
