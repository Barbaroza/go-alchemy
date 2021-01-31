package switcher

import (
	"fmt"
	"math/rand"
)

func Switcher() {
	//form1
	var input = 1
	switch {
	case input > 0:
		fmt.Printf("input > 0: %d \n", input)
		input -= 2
		//关键字 - 继续执行后续分支的代码
		fallthrough
	case input < 0:
		fmt.Printf("input < 0: %d \n", input)

	default:
		fmt.Printf("default: %d \n", input)
	}
	//form2
	switch input {
	case -1:
		fmt.Printf("-1: %d \n", input)
		input += 2
	case 1:
		fmt.Printf("1: %d \n", input)
	default:
		fmt.Printf("default: %d \n", input)
	}
	//form3
	switch result := rand.Int(); {
	case result < 0:
		fmt.Printf("random value : %d < 0 \n", result)
	case result > 0:
		fmt.Printf("random value : %d > 0 \n", result)
	default:
		fmt.Printf("random value : %d = 0 \n", result)
	}
}
