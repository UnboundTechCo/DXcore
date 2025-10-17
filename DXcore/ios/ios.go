package ios

import (
	"fmt"
)

type ProgressListener interface {
	OnProgress(msg string)
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

func GetFlowLine() string {
	fmt.Println("GetFlowLine")
	return ""
}

func Log(message string) {
	fmt.Println(message)
}
