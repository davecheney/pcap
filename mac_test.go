package pcap

import (
  "testing"
  )

const (
  EXPECTED = "00:26:f2:5b:ac:4d"
  )

func TestMacPrinting(t *testing.T) {
  mac := Mac([...]byte { 0x0, 0x26, 0xf2, 0x5b, 0xac, 0x4d })
  if mac.String() != EXPECTED {
    t.Fatalf("Expected %s, Actual %s", EXPECTED, mac.String())
  }
}
