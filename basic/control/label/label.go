package label

import "fmt"

func Label() {
	a := 1
	//goto TARGET // compile error
	b := 9
	b += a
	fmt.Printf("a is %v *** b is %v", a, b)
}
