package config

import (
	"github.com/metacubex/mihomo/adapter"
	"github.com/metacubex/mihomo/constant"
)

type Rules struct {
	Domain  []string `json:"domain,omitempty"`
	Process []string `json:"process,omitempty"`
}

type Options struct {
	Port  uint16         `json:"port"`
	Proxy map[string]any `json:"proxy"`
	Rules Rules          `json:"rules,omitempty"`

	Dialer constant.Proxy `json:"-"`
}

func (o *Options) ParseProxy() error {
	dialer, err := adapter.ParseProxy(o.Proxy)
	if err != nil {
		return err
	}
	o.Dialer = dialer
	return nil
}
