package service

import (
	"bufio"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/InWILL/MioSocks/engine"
	"github.com/InWILL/MioSocks/socks"
	"github.com/metacubex/mihomo/adapter"
	"github.com/metacubex/mihomo/constant"
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

func (m *MioEngine) DelayTest(proxy map[string]any) (int64, error) {
	metadata := &constant.Metadata{
		NetWork: constant.TCP,
		Host:    "clients3.google.com",
		DstPort: 80,
	}
	dialer, err := adapter.ParseProxy(proxy)
	if err != nil {
		fmt.Println("Parse proxy error:", err)
		return 0, err
	}
	ctx := context.Background()
	conn, err := dialer.DialContext(ctx, metadata)
	if err != nil {
		fmt.Println("DialContext error:", err)
		return 0, err
	}
	defer conn.Close()

	start := time.Now()

	// 发送 HTTP GET 请求
	req := "GET /generate_204 HTTP/1.1\r\n" +
		"Host: clients3.google.com\r\n" +
		"Connection: close\r\n\r\n"

	_, err = conn.Write([]byte(req))
	if err != nil {
		panic(err)
	}

	// 读取响应头
	reader := bufio.NewReader(conn)
	statusLine, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Println("响应状态行:", strings.TrimSpace(statusLine))
	delay := time.Since(start)

	fmt.Println(delay)
	return delay.Milliseconds(), nil
}
