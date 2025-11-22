package service

import (
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
	engine engine.EngineInterface
	socks5 socks.Socks5Interface
}

func NewService(options MioOptions) (MioService, error) {
	dialer, err := adapter.ParseProxy(options.Proxy)
	if err != nil {
		return nil, err
	}

	engine := engine.NewEngine(
		engine.EngineOptions{
			Dialer:  dialer,
			Process: options.Rules.Process,
		})

	socks5 := socks.NewSocks5(
		socks.Socks5Options{
			Port:   options.Port,
			Dialer: dialer,
		})

	service := &MioEngine{
		engine: engine,
		socks5: socks5,
	}
	return service, nil
}

func (m *MioEngine) Start() {
	m.engine.Start()
	m.socks5.Start()
	NewRestAPI(62334)
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
