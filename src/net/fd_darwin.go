// +build darwin

package net

import (
	"syscall"
	"os"
)

func (fd *netFD) writeToSCTP(p []byte, sinfo *syscall.SCTPSndInfo, sa syscall.Sockaddr) (length int, err error) {
	if err := fd.writeLock(); err != nil {
		return 0, err
	}
	defer fd.writeUnlock()
	if err := fd.pd.PrepareWrite(); err != nil {
		return 0, err
	}
	for {
//		err = SCTPSendV(fd.sysfd, p, 0, sa)
		length, err = syscall.SCTPSendMsg(fd.sysfd, p, sinfo, sa, 0)

		if err == syscall.EAGAIN {
			if err = fd.pd.WaitWrite(); err == nil {
				continue
			}
		}
		break
	}

	if _, ok := err.(syscall.Errno); ok {
		err = os.NewSyscallError("sctpsendv", err)
	}
	return
}