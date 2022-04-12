package operational_system

import "log"

func distributor() {
	if showLogs() {
		log.Println("Starting distributor")
	}

	for {
		if isOsStatusTurningOff() {
			if showLogs() {
				log.Println("Turning off distributor")
			}

			break
		}

		if fe.isEmpty() {
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
