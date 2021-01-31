package _map

import "fmt"

func MapDemo() {
	temp := make(map[int]int)
	var tempPtr = new(map[int]int)

	fmt.Printf("make map len: %d \n", len(temp))
	fmt.Printf("make map memory: %p \n", temp)
	fmt.Printf("make map value: %s \n", temp)

	fmt.Printf("new map memory: %p \n", &tempPtr)
	fmt.Printf("new map value: %s \n", *tempPtr)

	for i := 0; i < 5; i++ {
		temp[i] = i
	}
	for k, _ := range temp {
		if k == 2 {
			delete(temp, 2)
		}
	}
	fmt.Printf("make map len: %d \n", len(temp))

}

func mapSliceInitDemo() {
	mapSlice := make([]map[int]int, 5, 10)
	for i := range mapSlice {
		mapSlice[i] = make(map[int]int, 1)
		mapSlice[i][1] = 1
	}
}

func sortMapDemo() {

}
