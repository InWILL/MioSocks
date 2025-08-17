package engine

import (
	"github.com/metacubex/mihomo/constant"
)

type EngineInterface interface {
	Start()
	Close()
	UpdateProxy(dialer constant.Proxy)
	GetUpStream() int64
	GetDownStream() int64
}
