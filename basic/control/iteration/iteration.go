package iteration

import "fmt"

func Iterate() {
	//base
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}
	//条件
	var index int = 0
	for index < 0 {
		index--
	}

	for {
		break
	}

}

func Range() {
	str := "practice makes better"
	for index, element := range str {
		fmt.Printf("index : %d ,element: %c", index, element)
	}

}

func Demo1() {

	s := ""
	for s != "aaaaa" {
		fmt.Println("Value of s: \n", s)
		s = s + "a"
	}

}
