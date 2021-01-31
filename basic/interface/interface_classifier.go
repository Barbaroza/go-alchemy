package _interface

import "fmt"

func Classifier(items ...interface{}) {

	for _, t := range items {
		switch t.(type) {
		case person:
			fmt.Printf("person")
		case machine:
			fmt.Printf("machine")
		case animal:
			fmt.Printf("animal")
		case nil:
			fmt.Printf("nil")
		default:
			fmt.Printf("default")

		}
	}
}
