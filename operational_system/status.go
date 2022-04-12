package operational_system

func setOsStatusTurningOn() {
	osStatus = osStatusTurningOn
}

func isosStatusTurningOn() bool {
	return osStatus == osStatusTurningOn
}

func setOsStatusOn() {
	osStatus = osStatusOn
}

func isOsStatusOn() bool {
	return osStatus == osStatusOn
}

func setOsStatusTurningOff() {
	osStatus = osStatusTurningOff
}

func isOsStatusTurningOff() bool {
	return osStatus == osStatusTurningOff
}

func setOsStatusOff() {
	osStatus = osStatusOff
}

func isOsStatusOff() bool {
	return osStatus == osStatusOff
}
