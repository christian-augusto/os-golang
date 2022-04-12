package operational_system

type process struct {
	/*
		Real Time: 0
		User: 1,2,3
	*/
	Priority uint8

	/*
		Seconds
	*/
	ProcessingTimeSeconds int

	/*
		Kb
	*/
	QtdMemory int

	Id                  int
	ArrivalTime         int
	QtdPrinters         int
	QtdScanners         int
	QtdModems           int
	QtdCds              int
	ProcessedTime       int
	CurrentProcessQueue string
}

func newProcess(
	priority uint8,
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
		Priority:              priority,
		ProcessingTimeSeconds: processingTimeSeconds,
		QtdMemory:             qtdMemory,
		Id:                    id,
		ArrivalTime:           arrivalTime,
		QtdPrinters:           qtdPrinters,
		QtdScanners:           qtdScanners,
		QtdModems:             qtdModems,
		QtdCds:                qtdCds,
		ProcessedTime:         processedTime,
		CurrentProcessQueue:   currentProcessQueue,
	}
}

func processCreator(
	priority uint8,
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
