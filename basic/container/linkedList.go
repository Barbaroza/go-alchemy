package container

import "container/list"

func linkedListDemo() {
	linkedList := list.New()
	linkedList.PushFront(1)
	linkedList.PushFront(2)
	linkedList.PushFront(3)
}
