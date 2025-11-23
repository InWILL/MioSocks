package socks

import (
	"fmt"
	"log"
	"net"

	"github.com/metacubex/mihomo/constant"
)

type Socks5 struct {
	Port     uint16
	Proxy    constant.Proxy
	isClosed bool
	listener net.Listener
}

type Socks5Options struct {
	Port uint16
	//AllowLAN bool
	Dialer constant.Proxy
}

func NewSocks5(options Socks5Options) *Socks5 {
	return &Socks5{
		Port:     options.Port,
		Proxy:    options.Dialer,
		isClosed: false,
	}
}

func (e *Socks5) UpdateSocks5(options Socks5Options) error {
	if e.Port != options.Port {
		e.UpdatePort(options.Port)
	}
	if e.Proxy != options.Dialer {
		e.UpdateProxy(options.Dialer)
	}
	return nil
}

func (e *Socks5) Start() {
	go func() {
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
	}()
}

func (e *Socks5) Close() {
	e.isClosed = true
	e.listener.Close()
}

func (e *Socks5) UpdatePort(port uint16) {
	e.Port = port
	e.Close()
}

func (e *Socks5) UpdateProxy(proxy constant.Proxy) {
	e.Proxy = proxy
}
