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

func (pcap *PcapFile) ReadPacket() (*PacketHeader, []byte, os.Error) {
	var hdr = new(PacketHeader)
	if err := binary.Read(pcap.ReadCloser, binary.LittleEndian, hdr) ; err != nil {
		return nil, nil, err
	}
	data := make([]byte, hdr.Caplen)
	_, err := pcap.ReadCloser.Read(data)	
	return hdr, data, err
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
