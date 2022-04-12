package operational_system

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
