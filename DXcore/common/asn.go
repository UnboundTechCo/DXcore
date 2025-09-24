package common

import (
	"strings"
)

func getAsnName() string {
	return strings.ToLower(state.allProviderKey)
}

func SetAsnName() {
	asnName := getAsnName()
	state.asnName = asnName
}
