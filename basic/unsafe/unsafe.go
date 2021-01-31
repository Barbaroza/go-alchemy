package unsafe

import (
	"fmt"
	unsafe2 "unsafe"
)

func UnSafeDemo() {
	var i = uint32(2)
	sizeof := unsafe2.Sizeof(i)
	fmt.Printf("bytes: %d", sizeof)
}
