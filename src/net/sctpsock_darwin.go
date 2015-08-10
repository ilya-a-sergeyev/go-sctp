package net
import (
	"time"
	"syscall"
)

// +build darwin

// SCTPAddr represents the address of a SCTP end point
type SCTPAddr struct {
	IP IP
	Port int
	Zone string
}

func (a *SCTPAddr) opAddr() Addr {
	if a == nil  {
		return nil
	}
	return a
}

type SCTPConn struct {
	conn
}

func newSCTPConn(fd *netFD) *SCTPConn {
	c := &SCTPConn{conn{fd}}
	setNoDelaySCTP(c.fd, true)
	return c
}

func DialSCTP(net string, laddr, raddr *SCTPAddr) (*SCTPConn, error) {
	switch net {
	case "sctp", "sctp4", "sctp6":
	default:
		return nil, &OpError{Op: "dial", Net: net, Source: laddr.opAddr(), Addr: raddr.opAddr(), Err: UnknownNetworkError(net)}
	}
	if raddr == nil {
		return nil, &OpError{Op: "dial", Net: net, Source: laddr.opAddr(), Addr: nil, Err: errMissingAddress}
	}
	return dialSCTP(net, laddr, raddr, noDeadline)
}

func dialSCTP(net string, laddr, raddr *SCTPAddr, deadline time.Time) (*SCTPConn, error) {
	// TODO syscall.SOCK_SEQPACKET can also be syscall.SOCK_STREAM
	fd, err := internetSocket(net, laddr, raddr, deadline, syscall.SOCK_SEQPACKET, 0, "dial")
	if err != nil {
		return nil, &OpError{Op: "dial", Net: net, Source: laddr.opAddr(), Addr: raddr.opAddr(), Err: err}
	}
	return newSCTPConn(fd), nil
}

func ResolveSCTPAddr(net, addr string) (*SCTPAddr, error) {
	switch net {
	case "sctp", "sctp4", "sctp6":
	case "":
		net = "sctp"
	default:
		return nil, UnknownNetworkError(net)
	}
	addrs, err := internetAddrList(net, addr, noDeadline)
	if err != nil {
		return nil, err
	}
	return addrs.first(isIPv4).(*SCTPAddr), nil
}

//
// Implement PacketConn interface
//

func (c *SCTPConn) ReadFrom(b []byte) (n int, addr Addr, err error) {
	return
}

func (c *SCTPConn) WriteTo(b []byte, addr Addr) (n int, err error) {
	if !c.ok() {
		return 0, syscall.EINVAL
	}
	a, ok := addr.(*SCTPAddr)
	if !ok {
		return 0, &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr, Err: syscall.EINVAL}
	}
	return c.WriteToSCTP(b, a)
}

func (c *SCTPConn) Close() error {
	return
}

func (c *SCTPConn) LocalAddr() Addr {
	return
}

func (c *SCTPConn) SetDeadline(t time.Time) error {
	return
}

func (c *SCTPConn) SetReadDeadline(t time.Time) error {
	return
}

func (c *SCTPConn) SetWriteDeadline(t time.Time) error {
	return
}

//
// SCTP specific implementations
//

func (a *SCTPAddr) sockaddr(family int) (syscall.Sockaddr, error) {
	if a == nil {
		return nil, nil
	}
	return ipToSockaddr(family, a.IP, a.Port, a.Zone)
}

func (c *SCTPConn) WriteToSCTP(b []byte, addr *SCTPAddr) (int, error) {
	if !c.ok() {
		return 0, syscall.EINVAL
	}
	if addr == nil {
		return 0, &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: nil, Err: errMissingAddress}
	}
	sa, err := addr.sockaddr(c.fd.family)
	if err != nil {
		return 0, &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr.opAddr(), Err: err}
	}
	n, err := c.fd.writeToSCTP(b, sa)
	if err != nil {
		err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr.opAddr(), Err: err}
	}
	return n, err
}

