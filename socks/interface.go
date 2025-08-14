package socks

import (
	"fmt"
	"log"
	"net"

	"github.com/metacubex/mihomo/constant"
)

type Socks5 interface {
	Start()
	Close()
	UpdatePort(port uint16)
	UpdateProxy(proxy constant.Proxy)
}

func NewSocks5(options Socks5Options) Socks5 {
	return &Socks5Engine{
		Port:     options.Port,
		Proxy:    options.Dialer,
		isClosed: false,
	}
}

func (e *Socks5Engine) Start() {
	addr := fmt.Sprintf(":%d", e.Port)

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	log.Printf("%s server: %s listening on %s", e.Proxy.Type(), e.Proxy.Name(), addr)

	e.listener = ln

	for {
		conn, err := ln.Accept()
		if err != nil {
			if e.isClosed {
				break
			}
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		e.HandleConnection(conn)
	}
}

func (e *Socks5Engine) Close() {
	e.isClosed = true
	e.listener.Close()
}

func (e *Socks5Engine) UpdatePort(port uint16) {
	e.Port = port
	e.Close()
}

func (e *Socks5Engine) UpdateProxy(proxy constant.Proxy) {
	e.Proxy = proxy
}
