package pcap

type Packet interface {
	Payload() []byte
}
