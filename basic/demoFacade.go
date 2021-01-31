package basic

import (
	array2 "basic/array"
	"basic/assignment"
	"basic/control/functioned"
	"basic/control/iteration"
	"basic/control/mutireturn"
	switcher2 "basic/control/switcher"
	_error "basic/error"
	_interface "basic/interface"
	_map "basic/map"
	"basic/method"
	"basic/pointer"
	"basic/reflect"
	slice2 "basic/slice"
	"basic/standard"
	"basic/structure"
	"basic/unsafe"
	"fmt"
)

func AssignmentDemo() {
	assignment.N()
	assignment.M()
	assignment.N()
}

func PointerDemo() {
	pointer.PrintMemoryInfo()
}
func PointerDemo2() {
	pointer.PrintMemoryInfo2()
}
func MutiReturn() {
	mutireturn.StringConvert()
}
func switcher() {
	switcher2.Switcher()
}
func forDemo() {
	iteration.Range()
}
func function() {
	//functioned.Min()
	//functioned.ParameterDemo()
	functioned.Anonymous()

}
func array() {
	array2.Declare()
}
func slice() {
	slice2.SliceDemo4()
}
func string() {
	slice2.StringDemo()
}
func mapDemo() {
	_map.MapDemo()
}
func unSafeDemo() {
	unsafe.UnSafeDemo()
}
func regexpDemo() {
	standard.RegexpDemo()
}

func structureDemo() {
	structure.StructureDemo()
}

func structureDemo2() {
	structure.StructureDemo()
	p := new(structure.Person)
	fmt.Printf("person: %v \n", p)
}

func reflectDemo() {
	structure.ReflectTagInfo()
}
func anonymousFieldDemo() {
	structure.ExtendsDemo()
}

func DuplicateField() {
	structure.DuplicateField()
}

func ReceiverDemo() {
	method.ReceiverDemo()

}

func DiffReceiverDemo() {
	method.DiffMethodDemo()
}

func ExtenderDemo() {
	method.ExtenderDemo()
}

func InterfaceDemo() {
	_interface.InterfaceDemo()
}

func InterfaceClassifier() {
	_interface.AssertType()
}

func InterfaceSets() {
	_interface.InterfaceSetsDemo()
}

func InterfaceSort() {
	_interface.ArraySortDemo()
}

func ReflectSetValue() {
	reflect.ReflectSetDemo()
}
func ReflectStructure() {
	reflect.ReflectStructureDemo()
}

func ErrorProcessor() {
	_error.ErrorProcessorDemo()
}
func PanicRecover() {
	_error.PanicRecoverDemo()
}
