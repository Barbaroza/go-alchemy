package pointer

import "fmt"

func PrintMemoryInfo() {
	var value = 1
	fmt.Printf("integer :%d ,location in memory :%p \n", value, &value)
	var tempInt *int
	tempInt = &value
	fmt.Printf("integer :%d ,location in memory :%p \n", *tempInt, tempInt)
	fmt.Printf("%v", *&value == value)
}

func PrintMemoryInfo2() {
	var s = "good"
	var p *string
	p = &s
	*p = "bad"
	fmt.Printf("string:%s , memory location : %p \n", s, &s)
	fmt.Printf("string:%s , memory location : %p \n", *p, p)
}



