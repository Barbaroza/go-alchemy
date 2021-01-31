package functioned

import "fmt"

var (
	param1 = 1
	array2 = [...]int{4, 5, 6}
)

func ParameterDemo() {
	t1(param1)
	fmt.Printf("t1(param1) = %d\n", param1)

	t2(&param1)
	fmt.Printf("t2(param1) = %d\n", param1)

	t3(array2)
	fmt.Printf("t3(array2) = %+v\n", array2)

	t4(array2[0])
	fmt.Printf("t3(array2) = %+v\n", array2)

	t5(array2[:])
	fmt.Printf("t3(array2) = %+v\n", array2[:])

	func1 := f1
	t6(param1, func1)
	fmt.Printf("t6(param1, f1) = %d\n", param1)
}

// 1.变量（值拷贝）
func t1(tmp int) {
	tmp++
}

// 2.指针（地址拷贝）
func t2(tmp *int) {
	*tmp++
}

// 3.数组（值拷贝）
func t3(arr [3]int) {
	arr[0] = 1
}

// 4.数组元素(值拷贝)
func t4(tmp int) {
	tmp++
}

// 5.slice(地址拷贝)
func t5(s1 []int) {
	s1[0]++
}

// 6.函数（直接调用）
func t6(a int, f1 func(*int)) {
	f1(&a)
	fmt.Printf("f1(&a) = %d\n", a)
}

func f1(a *int) {
	*a++
}
