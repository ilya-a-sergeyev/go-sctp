// +build darwin

package net
import (
	"os"
	"syscall"
)

func setNoDelaySCTP(fd *netFD, noDelay bool) error {
	if err := fd.incref(); err != nil {
		return err
	}
	defer fd.decref()
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(fd.sysfd, syscall.IPPROTO_SCTP, syscall.SCTP_NODELAY, boolint(noDelay)))
}
