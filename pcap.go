package pcap

import (
	"io"
	"os"
)

type PacketReader interface {
	ReadPacket() (*Capture, os.Error)
	io.Closer
}
