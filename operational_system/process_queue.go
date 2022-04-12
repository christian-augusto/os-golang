package operational_system

import (
	"container/list"
)

type processQueue struct {
	queue *list.List
}

func newProcessQueue() *processQueue {
	return &processQueue{
		queue: list.New(),
	}
}

func (p *processQueue) isEmpty() bool {
	return p.size() == 0
}

func (p *processQueue) size() int {
	return p.queue.Len()
}

func (p *processQueue) insert(process *process) {
	p.queue.PushBack(process)
}

func (p *processQueue) remove() *process {
	firstProcess := p.queue.Remove(p.queue.Front())

	return firstProcess.(*process)
}
