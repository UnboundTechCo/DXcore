package android

import (
	"fmt"

	"github.com/unboundtech/DXcore-empty/DXcore/common"
)

type AndroidWarpClient struct {
	*common.WarpClient
}
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
	common.StartVPN(cacheDir, flowLine, pattern)
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
	common.Log(message)
}

func SetProgressListener(l ProgressListener) {
	common.SetProgressListener(l)
}
