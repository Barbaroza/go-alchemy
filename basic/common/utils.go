package common

import (
	"fmt"
	"runtime"
)

func Printf(formatter string, values ...interface{}) {
	fmt.Printf(formatter, values...)
}
func Println(values ...interface{}) {
	fmt.Println(values)
}
func Where() func() {
	return func() {
		_, file, line, _ := runtime.Caller(1)
		Printf("%s:%d\n", file, line)
	}
}
