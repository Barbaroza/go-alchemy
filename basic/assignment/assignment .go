package assignment

import "fmt"

var a = "G"
var b = "G"

func N() {
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(&b)
}

func M() {
	a := "O"
	b = "O"
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(&b)

}
