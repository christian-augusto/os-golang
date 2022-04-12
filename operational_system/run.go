package operational_system

import (
	"os-golang/utils"
)

func Run() {
	// memory positions 1.048.576, 1 Kb each one

	loadEnvironmentVars()

	start()

	stop()
}

func start() {
	fe = newProcessQueue()
	ftr = newProcessQueue()
	fu = newProcessQueue()
	fu2 = newProcessQueue()
	fu3 = newProcessQueue()

	asyncProcess.Add(1)

	go distributor()

	utils.WaitSeconds(10)

	setOsStatusTurningOff()

	asyncProcess.Wait()

	stop()
}

func stop() {
	fe = nil
	ftr = nil
	fu = nil
	fu2 = nil
	fu3 = nil

	setOsStatusOff()
}
