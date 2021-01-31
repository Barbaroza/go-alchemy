package _interface

func EmptyInterfaceConvert() {
	p := person(new(man))
	var obj Object = p


	obj.(person).speak()

	if person, ok := obj.(person); ok {
		person.speak()
	}
}
