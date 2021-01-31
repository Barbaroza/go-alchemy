package functioned

import "fmt"

func BlankIdentifier() {
	var res1, res2 int
	res1, res2, _ = threeValues()
	fmt.Print(res1 + res2)

	var min, max int
	min, max = minMax(78, 65)
	fmt.Printf("Minmium is: %d, Maximum is: %d\n", min, max)
}

func minMax(a int, b int) (min int, max int) {
	if a < b {
		min = a
		max = b
	} else { // a = b or a < b
		min = b
		max = a
	}
	return
}

func threeValues() (int, int, int) {
	return 1, 2, 3
}
