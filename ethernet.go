package pcap

import (
  "encoding/binary"
  )

func Decode(payload []byte) (Mac, Mac, uint16, []byte) {
  source, dest := make(Mac, 6), make(Mac, 6)
  copy(source,(payload[0:6]))
  copy(dest, (payload[6:13]))
  println(len(source))
  frame := binary.LittleEndian.Uint16(payload[13:15])
  return source, dest, frame, payload[15:]
}
