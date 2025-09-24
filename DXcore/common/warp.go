package common

import (
	"context"
)

type WarpClient struct {
	Ctx    context.Context
	Cancel context.CancelFunc
}

func (w *WarpClient) Stop() bool {
	return true
}

func (w *WarpClient) Start() error {
	return nil
}

func (w *WarpClient) IsRunning() bool {
	return false
}

func StartWithConfig(bindAddress, cacheDir string) bool {
	return true
}
