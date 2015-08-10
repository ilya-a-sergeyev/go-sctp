// +build darwin

package syscall

/*
#include <netinet/sctp.h>
*/

import "C"
import "unsafe"


func SCTPSendV(fd int, p []byte, flags int, to Sockaddr) (err error) {

	var iov Iovec
	if len(p) > 0 {
		iov.Base = (*byte)(unsafe.Pointer(&p[0]))
		iov.SetLen(len(p))
	}
	ptr, n, err := to.sockaddr()

	C.sctp_sendv(
		uintptr(fd),
		uintptr(unsafe.Pointer(&iov)),
		1,
		uintptr(unsafe.Pointer(to)),
		1,
		(void *) &sinfo,
		sizeof(struct sctp_sndinfo),
		SCTP_SENDV_SNDINFO,
		0)

	return
}
