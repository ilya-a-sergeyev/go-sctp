// +build darwin

package net
import (
	"syscall"
	"os"
)

func (fd *netFD) writeToSCTP(p []byte, sa syscall.Sockaddr) (n int, err error) {
	if err := fd.writeLock(); err != nil {
		return 0, err
	}
	defer fd.writeUnlock()
	if err := fd.pd.PrepareWrite(); err != nil {
		return 0, err
	}
	for {
		err = syscall.SCTPSendV(fd.sysfd, p, 0, sa)
		if err == syscall.EAGAIN {
			if err = fd.pd.WaitWrite(); err == nil {
				continue
			}
		}
		break
	}
	if err == nil {
		n = len(p)
	}
	if _, ok := err.(syscall.Errno); ok {
		err = os.NewSyscallError("sctpsendv", err)
	}
	return
}