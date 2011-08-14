package pcap

import (
	"testing"
	"os"
)

func TestReadPcapFile(t *testing.T) {
	pcap, err := Open("testdata/snmp.pcap")
	if err != nil {
		t.Error(err)
	}
	defer pcap.Close()
	fhdr, _ := pcap.(*PcapFile)
	t.Log(fhdr.hdr)
	
	for {
		capture, err := pcap.ReadCapture()
		if err != nil {
			if err == os.EOF {
				return
			}
			t.Fatal(err)
		}
		t.Log(capture)
		frame, err := ParseEthernetFrame(capture)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(frame)
		packet, err := ParseIP(frame)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(packet)
	}
}
