package method

import "fmt"

type Log struct {
	msg string
}

type TypeA struct {
	*Log
}

type TypeB struct {
	Log
}

func (receiver *TypeA) log() *Log {
	return receiver.Log
}

func (receiver *Log) append(msg string) {
	receiver.msg += msg + "\n"
}

func (receiver *Log) String() string {
	return receiver.msg
}

func (receiver TypeB) string() string {
	return fmt.Sprint(receiver.Log)
}

func DiffExtender() {
	typeA := new(TypeA)
	typeA.Log = new(Log)
	typeA.append("msg")
	fmt.Println(typeA.log())

	typeB := TypeB{Log{""}}
	typeB.append("msg")
	fmt.Println(typeB)
}
