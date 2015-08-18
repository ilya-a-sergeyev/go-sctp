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

func (fd *netFD) ReadFromSCTP(p []byte) (n int, oobn int, flags int, sa syscall.Sockaddr, rinfo *syscall.SCTPRcvInfo, err error) {
	if err = fd.readLock(); err != nil {
		return
	}
	defer fd.readUnlock()
	if err = fd.pd.PrepareRead(); err != nil {
		return
	}
	for {
//		(n int, oobn int, from Sockaddr, rinfo *SCTPRcvInfo, recvflags int, err error)
		n, oobn, sa, rinfo, flags, err = syscall.SCTPReceiveMessage(fd.sysfd, p, 0)

		if err != nil {
			n = 0
			if err == syscall.EAGAIN {
				if err = fd.pd.WaitRead(); err == nil {
					continue
				}
			}
		}
		err = fd.eofError(n, err)
		break
	}
	if _, ok := err.(syscall.Errno); ok {
		err = os.NewSyscallError("recvfrom", err)
	}
	return
}
