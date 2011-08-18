package pcap

type Frame interface {
	Header() []byte
	Payload() []byte
}
