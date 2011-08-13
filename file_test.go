package pcap

import (
	"testing"
)

func TestReadPcapFile(t *testing.T) {
	pcap, err := Open("testdata/snmp.pcap")
	if err != nil {
		t.Error(err)
	}
	defer pcap.Close()
	fhdr, _ := pcap.(*PcapFile)
	t.Log(fhdr.hdr)

	hdr, data, err := pcap.ReadPacket()
	if err != nil {
		t.Error(err)
	}
	if hdr.Caplen != uint32(len(data)) {
		t.Errorf("Expecting %d bytes by received %d", hdr.Caplen, len(data))
	}	
}
