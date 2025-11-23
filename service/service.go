package service

import (
	"github.com/InWILL/MioSocks/engine"
	"github.com/InWILL/MioSocks/socks"
	"github.com/metacubex/mihomo/adapter"
)

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
	options *MioOptions
	engine  *engine.Engine
	socks5  *socks.Socks5
}

func NewService(options MioOptions) (*MioEngine, error) {
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

	m := &MioEngine{
		options: &options,
		engine:  engine,
		socks5:  socks5,
	}
	return m, nil
}

func (m *MioEngine) UpdateService(options MioOptions) error {
	dialer, err := adapter.ParseProxy(options.Proxy)
	if err != nil {
		return err
	}

	m.socks5.UpdateSocks5(socks.Socks5Options{
		Port:   options.Port,
		Dialer: dialer,
	})
	m.engine.UpdateEngine(engine.EngineOptions{
		Dialer:  dialer,
		Process: options.Rules.Process,
	})

	m.options = &options
	return nil
}

func (m *MioEngine) Start() {
	m.engine.Start()
	m.socks5.Start()
	m.NewRestAPI(62334)
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
