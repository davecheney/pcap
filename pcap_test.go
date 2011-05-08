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
  _, err = pcap.ReadPacket() 
  if err != nil {
    t.Fatal(err)
  }
}

