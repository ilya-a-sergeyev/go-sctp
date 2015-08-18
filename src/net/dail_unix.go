// +build linux netbsd openbsd freebsd solaris

package net


// ListenPacket announces on the local network address laddr.
// The network net must be a packet-oriented network: "udp", "udp4",
// "udp6", "ip", "ip4", "ip6" or "unixgram".
// See Dial for the syntax of laddr.
func ListenPacket(net, laddr string) (PacketConn, error) {
	addrs, err := resolveAddrList("listen", net, laddr, noDeadline)
	if err != nil {
		return nil, &OpError{Op: "listen", Net: net, Source: nil, Addr: nil, Err: err}
	}
	var l PacketConn
	switch la := addrs.first(isIPv4).(type) {
	case *UDPAddr:
		l, err = ListenUDP(net, la)
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
