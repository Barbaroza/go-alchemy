package testing

import (
	"basic/coroutines"
	"testing"
)

func TestGoroutinesDemo(t *testing.T) {
	coroutines.GoroutinesDemo(2)
}

func TestGoroutinesDemo2(t *testing.T) {
	coroutines.GoroutinesDemo2()
}

func TestConsumeProvider(t *testing.T) {
	coroutines.ConsumeProvider()
}

func TestPrimeDemo(t *testing.T) {
	coroutines.PrimeDemo()
}
func TestPrimeDemo2(t *testing.T) {
	coroutines.PrimeDemo2()
}
func TestGoroutinesSelectDemo(t *testing.T) {
	coroutines.GoroutinesSelectDemo()
}

func TestGoroutinesRecoverDemo(t *testing.T) {
	coroutines.GoroutinesRecoverDemo()
}

func TestLazyEvaluatorDemo(t *testing.T) {
	coroutines.LazyEvaluatorDemo()
}
func TestMultiCoresDemo(t *testing.T) {
	coroutines.MultiCoresDemo()
}
