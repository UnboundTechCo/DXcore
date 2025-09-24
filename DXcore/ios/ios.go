package ios

import (
	"fmt"

	"github.com/unboundtech/fake-dxcore/DXcore/common"
)

type IosWarpClient struct {
	*common.WarpClient
}

type ProgressListener interface {
	OnProgress(msg string)
}

func Stop() bool {
	return false
}

func MeasurePing() int {
	return 0
}

func GetFlag() string {
	return "xx"
}

func StartVPN(cacheDir, flowLine, pattern string) {
	common.StartVPN(cacheDir, flowLine, pattern)
}

func StopVPN() bool {
	return true
}

func SetAsnName() {
	fmt.Println("SetAsnName")
}

func SetTimeZone(timeDiff float32) bool {
	return true
}

func GetFlowLine() string {
	return ""
}
