package functioned

import (
	"basic/common"
)

/*
 *
 */
func Anonymous() {
	where := common.Where()
	for i := 0; i < 5; i++ {
		a := func(i int) {
			common.Printf("current int: %d\n", i)
			where()
		}
		a(i)
		common.Printf("- g is of type %T and has value %p \n", a, &a)
	}
}
