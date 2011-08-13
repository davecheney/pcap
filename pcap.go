package pcap

import (
	"io"
)

type PacketReader interface {
	//	ReadPacket() (*Capture, os.Error)
	io.Closer
}
