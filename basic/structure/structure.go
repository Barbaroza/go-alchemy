package structure

import "fmt"

type Person struct {
	age  int32
	name string
}
type person struct {
	age  int32
	name string
}

func modify(p *Person, age int32, name string) {
	p.age = age
	p.name = name
}

func StructureDemo() {
	person := new(Person)
	fmt.Printf("new Persion : %v \n", person)
	person.name = "madonna"
	person.age = 56
	fmt.Printf("visable assignment Persion %v \n", person)
	modify(person, 32, "ray")
	fmt.Printf("modify new Persion : %v \n", person)
	instance0 := getInstance0(23, "lee")
	fmt.Printf("invisable new Persion1 : %v \n", instance0)
	instance1 := getInstance1(52, "Nat")
	fmt.Printf("invisable new Persion1 : %v \n", instance1)
	instance2 := getInstance2(12, "Sam")
	fmt.Printf("invisable new Persion2 : %v \n", instance2)
}
func getInstance0(age int32, name string) *Person {
	person := new(Person)
	person.age = age
	person.name = name
	return person
}
func getInstance1(age int32, name string) *Person {
	return &Person{age: age, name: name}
}

func getInstance2(age int32, name string) *Person {
	var person Person
	person.name = name
	person.age = age
	return &person
}

//Factory 通过可见性规则 禁用new 使用工厂模式
func NewPerson(age int32, name string) *person {
	return &person{age: age, name: name}
}
