package pcap

import (
	"io"
	"os"
)

type PacketReader interface {
	ReadPacket() (*PacketHeader, []byte, os.Error)
	io.Closer
}
