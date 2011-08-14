package pcap

type Capture interface {
	Payload() []byte
}
