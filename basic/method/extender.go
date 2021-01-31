package method

import "fmt"

type P struct {
	a int
	b int
	Action
}
type Action interface {
	eat()
}
type S struct {
	P
	Action
}

func (receiver *P) sum() int {
	return receiver.a + receiver.b
}
func (receiver *S) eat() {
	fmt.Printf("s.eat()\n")
}
func (receiver *P) eat() {
	fmt.Printf("p.eat()\n")
}

func ExtenderDemo() {
	s := new(S)
	s.a = 1
	s.b = 2
	sum := s.sum()
	s.eat()
	p := s.P
	p.eat()
	fmt.Printf("s.sum() result: %d \n", sum)
}
