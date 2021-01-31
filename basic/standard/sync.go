package standard

import (
	"container/list"
	"sync"
)

type Queue struct {
	lock  sync.Mutex
	queue list.List
}

func Update(queue *Queue) {
	queue.lock.Lock()
	queue.queue.PushFront(1)
	defer queue.lock.Unlock()
}
