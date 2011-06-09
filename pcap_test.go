package pcap

import (
	"testing"
)

func TestOnePacket(t *testing.T) {
	pcap, err := Open()
	if err != nil {
		t.Fatal(err)
	}
	defer pcap.Close()
        frame, err := pcap.ReadPacket()
	if err != nil {
		t.Fatal(err)
	}
        src, dest, f, payload := Decode(frame.payload)
        t.Logf("%s %s %#04x %#v", src, dest, f, payload)
}
