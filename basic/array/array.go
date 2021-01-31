package array

import "basic/common"

func Declare() {
	a := [3]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	c := [3]int{0: 1, 2: 3}
	assignment(a, &b, &c)
	common.Println(a)
	common.Println(b)
	common.Println(c)
}

func assignment(a [3]int, p1 *[3]int, p2 *[3]int) {
	a[0] = -1
	(*p1)[0] = -1
	t := *p2
	t[0] = -1
}
