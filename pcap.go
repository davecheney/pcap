package pcap

import (
	"io"
	"os"
)

type Packet interface {
	Data() []byte
}

type PacketReader interface {
	ReadCapture() (Capture, os.Error)
	io.Closer
}
