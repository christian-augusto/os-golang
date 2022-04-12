package operational_system

type process struct {
	/*
		REAL_TIME
		USER
		USER_2
		USER_3
	*/
	priority string

	/*
		Seconds
	*/
	processingTimeSeconds int

	/*
		Kb
	*/
	qtdMemory int

	id                  int
	arrivalTime         int
	qtdPrinters         int
	qtdScanners         int
	qtdModems           int
	qtdCds              int
	processedTime       int
	currentProcessQueue string
}

func newProcess(
	priority string,
	processingTimeSeconds int,
	qtdMemory int,
	id int,
	arrivalTime int,
	qtdPrinters int,
	qtdScanners int,
	qtdModems int,
	qtdCds int,
	processedTime int,
	currentProcessQueue string,
) *process {
	return &process{
		priority:              priority,
		processingTimeSeconds: processingTimeSeconds,
		qtdMemory:             qtdMemory,
		id:                    id,
		arrivalTime:           arrivalTime,
		qtdPrinters:           qtdPrinters,
		qtdScanners:           qtdScanners,
		qtdModems:             qtdModems,
		qtdCds:                qtdCds,
		processedTime:         processedTime,
		currentProcessQueue:   currentProcessQueue,
	}
}

func (p *process) isRealTime() bool {
	return p.priority == "REAL_TIME"
}

func (p *process) isUser() bool {
	return p.priority == userPriority
}

func (p *process) isUser2() bool {
	return p.priority == user2Priority
}

func (p *process) isUser3() bool {
	return p.priority == user3Priority
}

func processCreator(
	priority string,
	processingTimeSeconds int,
	qtdMemory int,
	id int,
	arrivalTime int,
	qtdPrinters int,
	qtdScanners int,
	qtdModems int,
	qtdCds int,
	processedTime int,
	currentProcessQueue string,
) {
	process := newProcess(
		priority,
		processingTimeSeconds,
		qtdMemory,
		id,
		arrivalTime,
		qtdPrinters,
		qtdScanners,
		qtdModems,
		qtdCds,
		processedTime,
		currentProcessQueue,
	)

	fe.insert(process)
}
