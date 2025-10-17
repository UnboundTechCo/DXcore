package android

import (
	"fmt"
)

type ProgressListener interface {
	OnProgress(msg string)
}

func StartT2S(tunfd int, bindAddress string) {
	err := StartTun2socks(tunfd, bindAddress)
	if err != nil {
		fmt.Println(err)
	}
}

func StopT2S() {
	StopTun2socks()
}

func Stop() bool {
	fmt.Println("Stop")
	return false
}

func MeasurePing() int {
	fmt.Println("MeasurePing")
	return 0
}

func GetFlag() string {
	fmt.Println("GetFlag")
	return "xx"
}

func StartVPN(cacheDir, flowLine, pattern string) {
	fmt.Println("StartVPN")
}

func StopVPN() bool {
	fmt.Println("StopVPN")
	return true
}

func SetAsnName() {
	fmt.Println("SetAsnName")
}

func SetTimeZone(timeDiff float32) bool {
	fmt.Println("SetTimeZone")
	return true
}

func GetFlowLine(isTest bool) string {
	fmt.Println("GetFlowLine")
	return ""
}

func Log(message string) {
	fmt.Println(message)
}

func SetProgressListener(l ProgressListener) {
}
