package pcap

import (
	"encoding/binary"
	"os"
)

func Decode(payload []byte) (Mac, Mac, uint16, []byte) {
	source, dest := make(Mac, 6), make(Mac, 6)
	copy(source, (payload[0:6]))
	copy(dest, (payload[6:13]))
	println(len(source))
	frame := binary.LittleEndian.Uint16(payload[13:15])
	return source, dest, frame, payload[15:]
}

type ethernetHeader struct {
	dest []byte
	src []byte
	vlan int
}

type ethernetFrame struct {
	header ethernetHeader
	payload []byte
}

func (e *ethernetFrame) Length() int {
	return len(e.payload)
}

func (e *ethernetFrame) Payload() (Packet, os.Error) {
	return parseUDP(e.payload)
}

func parseEthernet(data []byte) (Frame, os.Error) {
	
	return new(ethernetFrame), nil
}
