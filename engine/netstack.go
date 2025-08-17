package engine

import (
	"context"
	"io"
	"log"
	"net"
	"net/netip"

	"github.com/eycorsican/go-tun2socks/core"
	"github.com/metacubex/mihomo/constant"
)

var upstream int64
var downstream int64

type tcpHandler struct {
	constant.Proxy
}

func (e *Engine) NewStack() io.Writer {
	core.RegisterTCPConnHandler(NewTCPHandler(e.dialer))
	core.RegisterOutputFn(e.NetStack_Output)
	netstack := core.NewLWIPStack()
	return netstack
}

func NewTCPHandler(dialer constant.Proxy) core.TCPConnHandler {
	return &tcpHandler{
		Proxy: dialer,
	}
}

func (h *tcpHandler) Handle(conn net.Conn, target *net.TCPAddr) error {
	log.Printf("%s => %s:%d\n", conn.LocalAddr().String(), target.IP, target.Port)
	metadata := &constant.Metadata{
		NetWork: constant.TCP,
	}
	metadata.DstIP = netip.MustParseAddr(target.IP.String())
	metadata.DstPort = uint16(target.Port)

	ctx := context.Background()
	dstConn, err := h.DialContext(ctx, metadata)
	if err != nil {
		return err
	}

	go forward_upstream(dstConn, conn)
	go forward_downstream(conn, dstConn)

	return nil
}

func forward_upstream(dst, src net.Conn) {
	defer src.Close()
	defer dst.Close()
	size, _ := io.Copy(dst, src)
	upstream += size
}

func forward_downstream(dst, src net.Conn) {
	defer src.Close()
	defer dst.Close()
	size, _ := io.Copy(dst, src)
	downstream += size
}

func (e *Engine) UpdateProxy(dialer constant.Proxy) {
	core.RegisterTCPConnHandler(NewTCPHandler(dialer))
}

func (e *Engine) GetUpStream() int64 {
	return upstream
}

func (e *Engine) GetDownStream() int64 {
	return downstream
}
