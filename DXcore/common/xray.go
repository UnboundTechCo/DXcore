package common

import (
	"sync"
)

var xrayState struct {
	ClientMutex sync.Mutex
}

func StopXray() {
	xrayState.ClientMutex.Lock()
	defer xrayState.ClientMutex.Unlock()
}

func IsXrayRunning() bool {
	return true
}

func StartXray(link string) bool {
	return true
}
