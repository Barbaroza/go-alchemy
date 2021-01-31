package error

import "fmt"

func mockPanicRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("sub recover:%s\n", err)
			panic("quit by recover")
		}
	}()
	panic("quit")
}

func PanicRecoverDemo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("parent recover:%s\n", err)
		}
	}()
	fmt.Println("start")
	mockPanicRecover()
	//不会运行
	fmt.Println("end")

}
