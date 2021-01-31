package standard

import (
	"fmt"
	"regexp"
)

// 正则测试
func RegexpDemo() {
	a := "aosdjo123adasd"
	pet := "[0-9]"
	if match, _ := regexp.Match(pet, []byte(a)); match {
		fmt.Println("Match Found!")
	}
	re, _ := regexp.Compile(pet)
	str := re.ReplaceAllString(a, "-")
	fmt.Print(str)
}
