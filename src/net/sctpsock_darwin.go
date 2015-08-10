package net

// +build darwin

// SCTPAddr represents the address of a SCTP end point
type SCTPAddr struct {
	IP IP
	Port int
	Zone string
}

type SCTPConn struct {
	conn
}

func newSCTPConn(fd *netFD) *SCTPConn {
	c := &SCTPConn{conn{fd}}
	setNoDelaySCTP(c.fd, true)
	return c
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