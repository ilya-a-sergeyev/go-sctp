// +build darwin freebsd
package net

import (
	"context"
)

func ListenPacket(net, laddr string) (PacketConn, error) {
	addrs, err := DefaultResolver.resolveAddrList(context.Background(), "listen", net, laddr, nil)
	if err != nil {
		return nil, &OpError{Op: "listen", Net: net, Source: nil, Addr: nil, Err: err}
	}
	var l PacketConn
	switch la := addrs.first(isIPv4).(type) {
	case *UDPAddr:
		l, err = ListenUDP(net, la)
	case *SCTPAddr:
		l, err = ListenSCTP(net, la)
	case *IPAddr:
		l, err = ListenIP(net, la)
	case *UnixAddr:
		l, err = ListenUnixgram(net, la)
	default:
		return nil, &OpError{Op: "listen", Net: net, Source: nil, Addr: la, Err: &AddrError{Err: "unexpected address type", Addr: laddr}}
	}
	if err != nil {
		return nil, err // l is non-nil interface containing nil pointer
	}
	return l, nil
}
