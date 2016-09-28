# SCTP in Go

This branch has [SCTP](https://en.wikipedia.org/wiki/Stream_Control_Transmission_Protocol) functionality added to the network library.
I will try and keep this branch up to date with the master branch of the [official Go repository](https://github.com/golang/go)

## Go wild!
Needless to so say, SCTP in Go is experimental and should be used with caution.  

## Supported platforms 
For now this will only work on [FreeBSD](https://www.freebsd.org/) and [macOS/OSX](http://www.apple.com/uk/macos/sierra). 

FreeBSD comes with natiive support. On Mac OSX follow the instructions in the 
[official SCTP repository](https://github.com/sctplab/SCTP_NKE_ElCapitan) to install the driver.

Example server:

```golang
package main
import (
	"log"
	"os"
	"net"
)

func main() {

	saddr := "127.0.0.1:4242"
	addr, err := net.ResolveSCTPAddr("sctp", saddr)
	if err != nil {
		println(err)
		os.Exit(-1)
	}

	conn, err := net.DialSCTP("sctp", nil, addr)
	if err != nil {
		println("Error listening " + err.Error())
		os.Exit(-1)
	}

	defer conn.Close()

	var message = "hello"
	bmessage := []byte(message)

	_, err = conn.WriteTo(bmessage, addr)
	if err != nil {
		log.Printf("WriteTo error: %v", err)
	}

}
```

## Build instructions for SCTP in Go
These commands are based on the instructions [here](https://golang.org/doc/install/source).

	$ git clone https://github.com/cyberroadie/go-sctp
	$ cd go-sctp
	$ git checkout go-sctp
	$ cd src
	$ ./all.bash

## Test SCTP in Go
The [SCTP examples repository](https://github.com/cyberroadie/sctp-examples) contains working examples 
and instructions to test SCTP. It has TCP examples to compare with.

## Questions
Any questions drop me an email [ovanac01 at mail.bbk.ac.uk] or tweet #cyberroadie

Have fun! [Olivier Van Acker]



# The Go Programming Language

Go is an open source programming language that makes it easy to build simple,
reliable, and efficient software.

![Gopher image](doc/gopher/fiveyears.jpg)

For documentation about how to install and use Go,
visit https://golang.org/ or load doc/install-source.html
in your web browser.

Our canonical Git repository is located at https://go.googlesource.com/go.
There is a mirror of the repository at https://github.com/golang/go.

Go is the work of hundreds of contributors. We appreciate your help!

To contribute, please read the contribution guidelines:
	https://golang.org/doc/contribute.html

##### Note that we do not accept pull requests and that we use the issue tracker for bug reports and proposals only. Please ask questions on https://forum.golangbridge.org or https://groups.google.com/forum/#!forum/golang-nuts.

Unless otherwise noted, the Go source files are distributed
under the BSD-style license found in the LICENSE file.

--

## Binary Distribution Notes

If you have just untarred a binary Go distribution, you need to set
the environment variable $GOROOT to the full path of the go
directory (the one containing this file).  You can omit the
variable if you unpack it into /usr/local/go, or if you rebuild
from sources by running all.bash (see doc/install-source.html).
You should also add the Go binary directory $GOROOT/bin
to your shell's path.

For example, if you extracted the tar file into $HOME/go, you might
put the following in your .profile:

	export GOROOT=$HOME/go
	export PATH=$PATH:$GOROOT/bin

See https://golang.org/doc/install or doc/install.html for more details.
