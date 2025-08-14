package service

import (
	"github.com/InWILL/MioSocks/engine"
	"github.com/InWILL/MioSocks/socks"
	"github.com/metacubex/mihomo/adapter"
)

type MioService interface {
	Start()
	Close()
}

type Rules struct {
	Domain  []string `json:"domain,omitempty"`
	Process []string `json:"process,omitempty"`
}

type MioOptions struct {
	Port  uint16         `json:"port"`
	Proxy map[string]any `json:"proxy"`
	Rules Rules          `json:"rules,omitempty"`
}

type MioEngine struct {
	engine engine.Engine
	socks5 socks.Socks5
}

func NewService(options MioOptions) (MioService, error) {
	dialer, err := adapter.ParseProxy(options.Proxy)
	if err != nil {
		return nil, err
	}

	socks5 := socks.NewSocks5(
		socks.Socks5Options{
			Port:   options.Port,
			Dialer: dialer,
		})
	service := &MioEngine{
		engine: *engine.NewEngine(dialer),
		socks5: socks5,
	}
	return service, nil
}

func (m *MioEngine) Start() {
	m.engine.Start()
	m.socks5.Start()
}

func (m *MioEngine) Close() {

}
