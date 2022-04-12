package operational_system

var fe *processQueue
var ftr *processQueue
var fu *processQueue
var fu2 *processQueue
var fu3 *processQueue

func Run() {
	// memory positions 1.048.576, 1 Kb each one

	start()
}

func start() {
	fe = newProcessQueue()
	ftr = newProcessQueue()
	fu = newProcessQueue()
	fu2 = newProcessQueue()
	fu3 = newProcessQueue()
}
