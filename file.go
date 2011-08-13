package pcap

import (
	"fmt"
	"io"
	"os"
	"encoding/binary"
)

type PcapFile struct {
	io.ReadCloser
	hdr	FileHeader
}

func (pcap *PcapFile) readFileHeader() (os.Error) {
	return binary.Read(pcap.ReadCloser, binary.LittleEndian, &pcap.hdr) 
}

func (h FileHeader) String() string {
	return fmt.Sprintf("Magic: %x, Version: %d.%d, Snaplen: %d", h.Magic, h.Major, h.Minor, h.Snaplen)
}

func Open(file string) (PacketReader, os.Error) {
	r, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	pcap := &PcapFile{ io.ReadCloser: r }
	err = pcap.readFileHeader()
	if err != nil {
		return nil, err
	}
	return pcap, nil
}
