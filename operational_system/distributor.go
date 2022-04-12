package operational_system

import (
	"log"
)

func distributor() {
	defer asyncProcess.Done()

	if showLogs() {
		log.Println("Starting distributor")
	}

	for true {
		if isOsStatusTurningOff() {
			if showLogs() {
				log.Println("Turning off distributor")
			}

			break
		}

		if showLogs() {
			log.Println("Reading FE")
		}

		if fe.isEmpty() {
			if showLogs() {
				log.Println("Empty FE")
			}

			continue
		}

		if showLogs() {
			log.Println("distributing...")
		}

		process := fe.remove()

		if process.isRealTime() {
			ftr.insert(process)
		} else if process.isUser() {
			fu.insert(process)
		}
	}

	if showLogs() {
		log.Println("distributor off")
	}
}
