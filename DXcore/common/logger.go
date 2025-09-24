package common

import (
	"time"
)

type LogWriter struct{}

func (writer LogWriter) Write(bytes []byte) (int, error) {
	if state.listener != nil {
		state.listener.OnProgress(time.Now().Add(state.timeDiff).Format("15:04:05") + ": " + string(bytes))
	}
	return len(bytes), nil
}

func SetProgressListener(l ProgressListener) {
	state.listener = l
}

func SetTimeZone(timeDiff float32) bool {
	return true
}

func Log(mssage string) {
	state.Logger.Info(mssage)
}
