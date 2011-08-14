package pcap

import (
	"encoding/binary"
	"fmt"
	"os"
)

func ParseEthernetFrame(c Capture) (Frame, os.Error) {
	data := c.Payload()
	etherType := binary.LittleEndian.Uint16(data[13:15])
	//length := int(binary.LittleEndian.Uint16(data[15:17]))
	return &EthernetFrame{ 
		header: ethernetHeader{
			dest: data[0:6],
			src: data[6:13],
			ethertype: etherType,
		}, 
		payload: data[17:],
	 }, nil
}

type ethernetHeader struct {
	dest []byte
	src []byte
	vlan int
	ethertype uint16
}

type EthernetFrame struct {
	header ethernetHeader
	payload []byte
}

func (e *EthernetFrame) Payload() []byte {
	return e.payload
}

func (e *EthernetFrame) Ethertype() uint16 {
	return e.header.ethertype
}

func (e *EthernetFrame) String() string {
	return fmt.Sprintf("Ethernet: payload=%d ethertype=%x", len(e.payload), e.Ethertype())
}
