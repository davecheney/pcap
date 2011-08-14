package pcap

type Frame interface {
	Payload() []byte
}
