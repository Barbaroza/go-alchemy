package testing

import (
	"basic/structure"
	"testing"
)

func TestNewPerson(t *testing.T) {
	structure.NewPerson(16, "ray")
}

func BenchmarkNewPerson(b *testing.B) {
	for i := 0; i < 5; i++ {
		person := structure.NewPerson(16, "ray")
		b.Logf("%d : %v", i, person)
	}
}
