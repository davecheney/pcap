package pcap

import (
	"io"
	"os"
)

type PacketReader interface {
	ReadPacket() (Packet, os.Error)
	io.Closer
}
