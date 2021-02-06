package testing

import (
	"basic/coroutines"
	"testing"
)

func TestFuture(t *testing.T) {
	future := coroutines.NewFuture(new())
	get, err := future.Get()
	
}
