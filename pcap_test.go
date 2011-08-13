package pcap

import (
	"testing"
)

func TestReadPcapFile(t *testing.T) {
	pcap, err := Open("testdata/snmp.pcap")
	if err != nil {
		t.Error(err)
	}
	fhdr, _ := pcap.(*PcapFile)
	t.Log(fhdr.hdr)	
}
