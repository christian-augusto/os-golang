package utils

import "time"

func WaitSeconds(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}
