package coroutines

type Obj struct {
	fieldB int
	filedA int
	serial chan func()
}

func NewObj() *Obj {
	ptr := &Obj{serial: make(chan func())}

	go func() {
		f := <-ptr.serial
		f()
	}()

	return ptr
}

func (obj *Obj) setA(a int) {
	obj.serial <- func() {
		obj.filedA = a
	}
}

func (obj *Obj) A() int {
	fChan := make(chan int)
	obj.serial <- func() {
		fChan <- obj.filedA
	}

	return <-fChan
}
