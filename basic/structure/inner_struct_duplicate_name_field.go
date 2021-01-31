package structure

import "fmt"

type A struct {
	B
	C
	b int32
	c int
}
type B struct {
	b int32
	c int32
}
type C struct {
	b int32
	c int64
}

func DuplicateField() {
	a := A{B{2, 1}, C{2, 1}, 2, 1}
	fmt.Printf("%v", a)
	a.B.b = 3
	a.C.b = 3
	a.c = 6
	a.b = 1
}
