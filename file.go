package pcap

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type PcapFile struct {
	io.ReadCloser
	hdr FileHeader
}

func (pcap *PcapFile) readFileHeader() error {
	return binary.Read(pcap.ReadCloser, binary.LittleEndian, &pcap.hdr)
}

type capture struct {
	hdr     PacketHeader
	payload []byte
}

func (c *capture) String() string {
	return fmt.Sprintf("capture: payload=%d", c.hdr.Caplen)
}

func (c *capture) Payload() []byte {
	return c.payload
}

func (pcap *PcapFile) ReadPacket() (Packet, error) {
	var capture = new(capture)
	if err := binary.Read(pcap.ReadCloser, binary.LittleEndian, &capture.hdr); err != nil {
		return nil, err
	}
	capture.payload = make([]byte, capture.hdr.Caplen)
	_, err := pcap.ReadCloser.Read(capture.payload)
	return capture, err
}

func (h FileHeader) String() string {
	return fmt.Sprintf("Magic: %x, Version: %d.%d, Snaplen: %d", h.Magic, h.Major, h.Minor, h.Snaplen)
}

func Open(file string) (PacketReader, error) {
	r, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	pcap := &PcapFile{io.ReadCloser: r}
	err = pcap.readFileHeader()
	if err != nil {
		return nil, err
	}
	return pcap, nil
}
