package functioned

import "basic/common"

func deferDemo1() {

	for index := 0; index < 5; index++ {
		defer common.Println(index)
	}
}

func deferTrace() {

	defer common.Printf("leave %s \n", "deferTrace")
	common.Printf("in %s \n", "deferTrace")
}
