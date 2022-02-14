package models

type Process struct {
	/*
		Real Time 0 - User
		1,2,3
	*/
	Priority uint8

	/*
		Seconds
	*/
	ProcessingTimeSeconds uint64

	/*
		Kb
	*/
	QtdMemory uint64

	Id                    uint64
	ArrivalTime           uint64
	QtdPrinters           uint64
	QtdScanners           uint64
	QtdModems             uint64
	QtdCds                uint64
	ProcessedTime         uint64
	CurrentProcessQueueId uint8
}

func NewProcess(
	priority uint8,
	processingTimeSeconds uint64,
	qtdMemory uint64,
	id uint64,
	arrivalTime uint64,
	qtdPrinters uint64,
	qtdScanners uint64,
	qtdModems uint64,
	qtdCds uint64,
	processedTime uint64,
	currentProcessQueueId uint8,
) Process {
	return Process{
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
		CurrentProcessQueueId: currentProcessQueueId,
	}
}
