package operational_system

import (
	"log"
)

type cpu struct {
	id int
}

func newCpu(id int) *cpu {
	return &cpu{
		id: id,
	}
}

func (c *cpu) dispatcher() {
	if showLogs() {
		log.Printf("CPU %v - Starting dispatcher\n", c.id)
	}

	for {
		if isOsStatusTurningOff() {
			if showLogs() {
				log.Printf("CPU %v - Turning off dispatcher\n", c.id)
			}

			break
		}

		// read from ftr
		// read from fu
		// read from fu2
		// read from fu3
	}

	if showLogs() {
		log.Printf("CPU %v - dispatcher off\n", c.id)
	}
}
