package engine

import (
	"github.com/metacubex/mihomo/constant"
)

type EngineInterface interface {
	Start()
	Close()
	UpdateProxy(dialer constant.Proxy)
}

func (e *Engine) UpdateProxy(dialer constant.Proxy) {
	UpdateProxy(dialer)
}
