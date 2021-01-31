package _interface

import "sort"

type DataArray struct {
	data []int
}

func (d DataArray) Len() int {

	return len(d.data)
}

func (d DataArray) Less(i, j int) bool {

	return d.data[i] < d.data[j]
}

func (d DataArray) Swap(i, j int) {
	d.data[i], d.data[j] = d.data[j], d.data[i]
}

func ArraySortDemo() {
	a := new(DataArray)
	a.data = []int{1, 3, 4, 5, 6, 71, 12, 2}
	array := DataArray{[]int{1, 2, 123, 5, 123, 2, 5}}
	sort.Sort(array)
	sort.Sort(a)
}
