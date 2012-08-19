package pcap

import "fmt"

type IPPacket interface {
	Header() []byte
	Payload() []byte
}

type IPv4Packet struct {
	header  []byte
	payload []byte
}

func (ip *IPv4Packet) Header() []byte {
	return ip.header
}

func (ip *IPv4Packet) Payload() []byte {
	return ip.payload
}

func (ip *IPv4Packet) String() string {
	return fmt.Sprintf("IPv4: payload=%d", len(ip.payload))
}

func ParseIP(frame Frame) (IPPacket, error) {
	data := frame.Payload()
	ihl := (data[0] & 0xf) << 2
	return &IPv4Packet{
		header:  data[:ihl],
		payload: data[ihl:],
	}, nil
}
