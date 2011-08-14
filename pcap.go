package pcap

import (
	"io"
	"os"
)

type Capture interface {
	Payload() Frame
	Body() []byte
	
}

type Frame interface {
	
}

type PacketReader interface {
	ReadPacket() (Capture, os.Error)
	io.Closer
}
