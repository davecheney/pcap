package pcap

import (
	"os"
)

type udpPacket struct {
	data []byte
}

func (u *udpPacket) Data() []byte {
	return u.data
}

func parseUDP(data []byte) (Packet, os.Error) {
	return &udpPacket{data}, nil
}	
