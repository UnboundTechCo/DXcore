package android

import (
	"github.com/unboundtech/fake-dxcore/DXcore/common"
)

// WarpConfig holds the configuration for Warp

type AndroidWarpClient struct {
	*common.WarpClient
}
type ProgressListener interface {
	OnProgress(msg string)
}

func Stop() bool {
	return common.Stop()
}

func StopT2S() {
	common.GetState().ClientMutex.Lock()
	defer common.GetState().ClientMutex.Unlock()

	StopTun2socks()
}

func StartT2S(tunfd int, bindAddress string) bool {
	return true
}

func MeasurePing() int {
	return common.MeasurePing()
}

func GetClientFlag() string {
	return common.GetClientFlag()
}

func GetFlag() string {
	return common.GetFlag()
}

func StartVPN(cacheDir, flowLine, pattern string) {
	common.StartVPN(cacheDir, flowLine, pattern)
}

func StopVPN() bool {
	return common.StopVPN()
}

func SetProgressListener(l ProgressListener) {
	common.SetProgressListener(l)
}

func SetAsnName() {
	common.SetAsnName()
}

func SetTimeZone(timeDiff float32) bool {
	return common.SetTimeZone(timeDiff)
}

func GetFlowLine() string {
	return common.GetFlowLine()
}

func Log(mssage string) {
	common.Log(mssage)
}
