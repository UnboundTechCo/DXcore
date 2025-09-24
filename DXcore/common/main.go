package common

import (
	"context"
	"log/slog"
	"sync"
	"time"
)

type ProgressListener interface {
	OnProgress(msg string)
}

type AppState struct {
	DefaultClient  *WarpClient
	ClientMutex    sync.Mutex
	Logger         *slog.Logger
	CancelFunc     context.CancelFunc
	PinUrls        []string
	listener       ProgressListener
	timeDiff       time.Duration
	vpnCancel      context.CancelFunc
	allProviderKey string
	asnName        string
}

var state = newState()

func newState() *AppState {
	writer := &LogWriter{}
	handler := slog.NewTextHandler(writer, &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if (a.Key == slog.TimeKey || a.Key == slog.LevelKey) && len(groups) == 0 {
				return slog.Attr{}
			}
			return a
		},
	})
	logger := slog.New(handler)

	return &AppState{
		PinUrls:  []string{},
		Logger:   logger,
		timeDiff: 0,
	}
}

func GetState() *AppState {
	return state
}

func Stop() bool {
	success := false
	return success
}

func startVPN() bool {
	return true
}

func StartVPN(cacheDir, flowLine, pattern string) {
	go startVPN()

}

func StopVPN() bool {
	if state.vpnCancel != nil {
		state.vpnCancel()
		state.vpnCancel = nil
		state.Logger.Info("VPN stopping")
	}
	return Stop()
}
