package models

import (
	"container/list"
)

type ProcessQueue struct {
	queue *list.List
}

func NewProcessQueue() ProcessQueue {
	return ProcessQueue{
		queue: list.New(),
	}
}

func (p ProcessQueue) Insert(process *Process) {
	p.queue.PushBack(process)
}

func (p ProcessQueue) Remove() *Process {
	firstProcess := p.queue.Remove(p.queue.Front())

	return firstProcess.(*Process)
}
