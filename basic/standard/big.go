package standard

import (
	"fmt"
	"math"
	"math/big"
)

func BigDemo() *big.Int {
	newInt := big.NewInt(math.MaxInt64)
	div := newInt.Div(newInt, big.NewInt(1))
	newRat := big.NewRat(math.MaxInt64, math.MaxInt64)
	fmt.Printf("big Rat: %v", newRat)
	return div
}
