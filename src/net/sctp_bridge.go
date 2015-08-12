// +build cgo
// +build darwin, amd64

package net

/*
#cgo CFLAGS: -I/usr/lib/libsctp.dylib
#cgo LDFLAGS: -L. -lsctp
#include <netinet/sctp.h>
#include <netinet/sctp_uio.h>

*/
import "C"
import (
	"unsafe"
	"syscall"
)

func SCTPSendV(fd int, p []byte, flags int, to syscall.Sockaddr) (err error) {

	var iov *syscall.Iovec
	if len(p) > 0 {
		iov.Base = (*byte)(unsafe.Pointer(&p[0]))
		iov.SetLen(len(p))
	}
//	ptr, _, err := to.sockaddr()
//	ptr := syscall.GetAddressPointer(to)

	var sinfo syscall.SCTPSndInfo
	sinfo.Sid = 1
//	sinfo.Ppid = 424242

//	C.sctp_sendv(
//		C.int(fd),
//		iov,
//		1,
//		C.struct_sockaddr(ptr),
//		1,
//		unsafe.Pointer(&sinfo),
//		syscall.SizeofSCTPSndInfo,
//		syscall.SCTP_SENDV_SNDINFO,
//		0)

	return nil
}
