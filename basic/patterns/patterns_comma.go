package patterns

import "fmt"

type function func() (interface{}, error)

var m = make(map[int]int)

var ch = make(chan int)

//处理函数错误

func processFuncError(f function) (interface{}, error) {
	i, err := f()
	if err != nil {
		fmt.Printf("Error %s in ", err.Error())
	}
	return i, err
}

func processFuncError2(f function) (interface{}, error) {
	var i interface{}
	var err error
	if i, err = f(); err != nil {
		fmt.Printf("Error %s in ", err.Error())
	}
	return i, err
}

//map中key是否存在

func keyIsPresent() {
	if _, ok := m[1]; !ok {
	}
}

//）检测一个通道ch是否关闭
func chanCloseCheck() {
	for v := range ch {
		fmt.Println(v)
	}

	for {
		if v, open := <-ch; open {
			fmt.Println(v)
		} else {
			break
		}

	}
}
