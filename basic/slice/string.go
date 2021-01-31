package slice

import (
	"fmt"
	"unicode/utf8"
)

func StringDemo() {
	s := "你好！世界"
	fmt.Printf("string len :%d \n", len(s))
	runes := []rune(s)
	bytes := []byte(s)
	fmt.Printf("rune len :%d \n", len(runes))
	fmt.Printf("bytes len :%d \n", len(bytes))
	fmt.Printf("runeCountInString len: %d \n", utf8.RuneCountInString(s))

	t := s[0:1]
	fmt.Printf("t:%d", int8(t[0]))

	runes[0] = '大'
	newString := string(runes)
	fmt.Printf("newString:%s", newString)
}
