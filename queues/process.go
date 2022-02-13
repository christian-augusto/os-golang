package process

import (
	"container/list"
	"os-golang/models/process"
)

type ProcessQueue struct {
	queue *list.List
}

func New() ProcessQueue {
	return ProcessQueue{
		queue: list.New(),
	}
}

func (p ProcessQueue) Insert(process *process.Process) {
	p.queue.PushBack(process)
}

func (p ProcessQueue) Remove() *process.Process {
	firstProcess := p.queue.Remove(p.queue.Front())

	return firstProcess.(*process.Process)
}
