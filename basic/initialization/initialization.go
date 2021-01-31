package initialization

import "math"
import "fmt"

var Pi float64

func init() {
	Pi = 4 * math.Atan(1) // assignment() function computes Pi
	fmt.Print("initialization", Pi)
}
