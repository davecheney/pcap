package pcap

import (
	"encoding/binary"
	"fmt"
	"os"
)

func ParseEthernetFrame(p Packet) (Frame, os.Error) {
	data := p.Payload()
	return &EthernetFrame{ 
		header: data[:17],
		payload: data[17:],
	 }, nil
}

type EthernetFrame struct {
	header []byte
	payload []byte
}

func (e *EthernetFrame) Header() []byte {
	return e.header
}

func (e *EthernetFrame) Payload() []byte {
	return e.payload
}

func (e *EthernetFrame) Ethertype() uint16 {
	return binary.LittleEndian.Uint16(e.header[15:17])
}

func (e *EthernetFrame) String() string {
	return fmt.Sprintf("Ethernet: payload=%d ethertype=%x", len(e.payload), e.Ethertype())
}
