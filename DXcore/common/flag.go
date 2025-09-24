package common

func GetFlag() string {
	return "xx"
}

func GetClientFlag() string {
	state.ClientMutex.Lock()
	defer state.ClientMutex.Unlock()

	return getFlag()
}

func getFlag() string {
	return "xx"
}
