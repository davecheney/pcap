package pcap

import (
	"fmt"
	"os"
)

type IPPacket interface {
	Payload() []byte	
}

type IPv4Header struct {

}

type IPv4Packet struct {
	header IPv4Header
	payload []byte
}

func (ip *IPv4Packet) Payload() []byte {
	return ip.payload
}

func (ip *IPv4Packet) String() string {
	return fmt.Sprintf("IPv4: payload=%d", len(ip.payload))	
}

func ParseIP(frame Frame) (IPPacket, os.Error) {
	data := frame.Payload()
	ihl := (data[0] & 0xf) << 2
	return &IPv4Packet {
		payload: data[ihl:],
	}, nil
}

