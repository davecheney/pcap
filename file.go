package pcap

import (
	"fmt"
	"io"
	"os"
	"encoding/binary"
)

type PcapFile struct {
	io.ReadCloser
	hdr FileHeader
}

func (pcap *PcapFile) readFileHeader() os.Error {
	return binary.Read(pcap.ReadCloser, binary.LittleEndian, &pcap.hdr)
}

type capture struct {
	hdr PacketHeader
	data	[]byte
}

func (c *capture) Payload() Frame {
	return nil
}

func (c *capture) Body() []byte {
	return c.data
}

func (pcap *PcapFile) ReadPacket() (Capture, os.Error) {
	var capture = new(capture)
	if err := binary.Read(pcap.ReadCloser, binary.LittleEndian, &capture.hdr) ; err != nil {
		return nil, err
	}
	capture.data = make([]byte, capture.hdr.Caplen)
	_, err := pcap.ReadCloser.Read(capture.data)
	return capture, err
}

func (h FileHeader) String() string {
	return fmt.Sprintf("Magic: %x, Version: %d.%d, Snaplen: %d", h.Magic, h.Major, h.Minor, h.Snaplen)
}

func Open(file string) (PacketReader, os.Error) {
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
