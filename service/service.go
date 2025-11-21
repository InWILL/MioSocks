package service

import (
	"github.com/InWILL/MioSocks/config"
	"github.com/InWILL/MioSocks/engine"
	"github.com/InWILL/MioSocks/socks"
	"github.com/metacubex/mihomo/adapter"
)

type MioService interface {
	Start()
	Close()
	UpdateProxy(Proxy map[string]any) error
	GetUpStream() int64
	GetDownStream() int64
}

type MioEngine struct {
	socks5 socks.Socks5Interface
	engine engine.EngineInterface
}

func NewService(options config.Options) (MioService, error) {
	err := options.ParseProxy()
	if err != nil {
		return nil, err
	}

	socks5 := socks.NewSocks5(options)
	engine := engine.NewEngine(options)
	service := &MioEngine{
		socks5: socks5,
		engine: engine,
	}
	return service, nil
}

func (m *MioEngine) Start() {
	m.engine.Start()
	m.socks5.Start()
}

func (m *MioEngine) Close() {

}

func (m *MioEngine) UpdateProxy(proxy map[string]any) error {
	dialer, err := adapter.ParseProxy(proxy)
	if err != nil {
		return err
	}

	m.engine.UpdateProxy(dialer)
	m.socks5.UpdateProxy(dialer)

	return nil
}

func (m *MioEngine) GetUpStream() int64 {
	return m.engine.GetUpStream() + m.socks5.GetUpStream()
}

func (m *MioEngine) GetDownStream() int64 {
	return m.engine.GetDownStream() + m.socks5.GetDownStream()
}
