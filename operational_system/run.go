package operational_system

import (
	"os-golang/utils"
	"sync"

	"github.com/spf13/viper"
)

func Run() {
	// memory positions 1.048.576, 1 Kb each one

	loadEnvironmentVars()

	start()

	stop()
}

func start() error {
	var err error

	setOsStatusTurningOn()
	fe = newProcessQueue()
	ftr = newProcessQueue()
	fu = newProcessQueue()
	fu2 = newProcessQueue()
	fu3 = newProcessQueue()
	cpusCount = viper.GetInt(cpusCountEnvName)
	asyncProcess = new(sync.WaitGroup)
	setOsStatusOn()

	asyncProcess.Add(1)
	go func() {
		distributor()

		asyncProcess.Done()
	}()

	cpus = make([]cpu, cpusCount)

	for i := 0; i < cpusCount; i++ {
		cpus[i] = *newCpu(i)

		asyncProcess.Add(1)
		go func(index int) {
			cpus[index].dispatcher()

			asyncProcess.Done()
		}(i)
	}

	utils.WaitSeconds(10)

	setOsStatusTurningOff()

	asyncProcess.Wait()

	stop()

	return err
}

func stop() {
	fe = nil
	ftr = nil
	fu = nil
	fu2 = nil
	fu3 = nil
	cpus = nil

	setOsStatusOff()
}
