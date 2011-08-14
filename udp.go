package pcap

import (
	"os"
)

type UDPHeader struct {

}

type Datagram struct {
	header UDPHeader
	payload	[]byte
}

func ParseUDP(ip IPPacket) (*Datagram, os.Error) {
	data := ip.Payload()
	return &Datagram {
		payload: data[20:],
	}, nil

}
