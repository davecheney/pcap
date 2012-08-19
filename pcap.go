package pcap

import "io"

type PacketReader interface {
	ReadPacket() (Packet, error)
	io.Closer
}
