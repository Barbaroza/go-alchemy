package functioned

import (
	"fmt"
)

func TypeCheck(values ...interface{}) {
	for _, value := range values {
		switch currentType := value.(type) {
		case int:
			fmt.Printf("%s", string(currentType))
		case float32:
		case bool:
		default:
		}
	}
}
