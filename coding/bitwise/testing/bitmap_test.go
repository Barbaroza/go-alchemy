package testing

import (
	"coding/bitwise"
	"fmt"
	"testing"
)

func TestBitMap(t *testing.T) {
	bitMap := bitwise.NewBitMap(546)
	fmt.Println(bitMap.IsExist(9))
	bitMap.Set(9)
	fmt.Println(bitMap.IsExist(9))
	bitMap.Remove(9)
	fmt.Println(bitMap.IsExist(9))

	bitMap.Set(1)
	bitMap.Set(5)
	bitMap.Set(546)
	bitMap.Set(465)
	bitMap.Set(23)

	fmt.Println(bitMap.Reverse())

}
