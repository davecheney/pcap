package pcap

import (
	"io"
	"os"
)

type Capture interface {
	Frame() (Frame, os.Error)
	Body() []byte
	
}

type Frame interface {
	Payload() (Packet, os.Error)
}

type Packet interface {
	Data() []byte
}

type PacketReader interface {
	ReadCapture() (Capture, os.Error)
	io.Closer
}
