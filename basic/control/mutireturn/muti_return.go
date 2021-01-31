package mutireturn

import (
	"fmt"
	"strconv"
)

func StringConvert() {
	var origin string = "ABC"
	atoi, err := strconv.Atoi(origin)
	if err != nil {
		fmt.Printf("%s is not int , error message : %s", origin, err.Error())
		return
	}
	fmt.Printf("%d", atoi)
}
