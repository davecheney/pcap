package pcap

type UDPHeader struct {
}

type Datagram struct {
	header  UDPHeader
	payload []byte
}

func ParseUDP(ip IPPacket) (*Datagram, error) {
	data := ip.Payload()
	return &Datagram{
		payload: data[20:],
	}, nil

}
