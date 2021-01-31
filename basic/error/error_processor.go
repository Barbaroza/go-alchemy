package error

import (
	"errors"
	"fmt"
)

var InvalidParamsError error = errors.New("invalid param nil")

func ValidateParams(params ...interface{}) error {
	if params == nil {
		return InvalidParamsError
	}
	for _, param := range params {
		if param == nil {
			return InvalidParamsError
		}
	}
	return nil
}

func ErrorProcessorDemo() {
	if err := ValidateParams(nil); err != nil {
		fmt.Printf("Error: %s", err)
	}
}
