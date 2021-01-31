package slice

import "fmt"

func SliceDemo() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5] // item at index 5 not included!

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

}

func SliceDemo2() {
	var temp = [5]int{1, 2, 3, 4, 5}
	a := func(slice []int) int {
		res := 0
		for i, _ := range slice {
			res = +slice[i]
		}
		return res
	}
	a(temp[:])
}

func SliceDemo3() {
	var slice []int = make([]int, 10, 10)
	for idx, value := range slice {
		slice[idx] = 2
		value++
	}
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice[i])
	}

	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2
	}

	for i := 0; i < len(items); i++ {
		fmt.Printf("items at %d is %d\n", i, items[i])
	}

}
func SliceDemo4() {
	t := make([]int, 5)
	for i := 0; i < len(t); i++ {
		fmt.Printf("Slice at %d is %d\n", i, t[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(t))
	fmt.Printf("The capacity of slice1 is %d\n", cap(t))

	reSize(t, 5)

	for i := 0; i < len(t); i++ {
		fmt.Printf("Slice at %d is %d\n", i, t[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(t))
	fmt.Printf("The capacity of slice1 is %d\n", cap(t))

	t = AppendByte(t, 2, 3, 4)
	for i := 0; i < len(t); i++ {
		fmt.Printf("Slice at %d is %d\n", i, t[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(t))
	fmt.Printf("The capacity of slice1 is %d\n", cap(t))

}

func reSize(t []int, factor int) {
	its := make([]int, factor*cap(t))
	copy(its, t)
	t = its
}

func AppendByte(slice []int, data ...int) []int {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]int, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
