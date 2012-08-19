package pcap

//#define _LARGEFILE_SOURCE
//#define _LARGEFILE64_SOURCE
//#define _FILE_OFFSET_BITS 64
//#define _GNU_SOURCE
//
//#include <sys/time.h>
//#include <sys/times.h>
//#include <sys/timex.h>
//#include <pcap/pcap.h>
//
//typedef struct pcap_file_header $FileHeader;
//typedef struct pcap_pkthdr $PacketHeader;
import "C"

const (
	Magic        = 0xa1b2c3d4
	MagicSwapped = 0xd4c3b2a1
)

type FileHeader struct {
	Magic uint32
	Major uint16
	Minor uint16
	Thiszone int32
	Sigfigs uint32
	Snaplen uint32
	Network uint32
}

type PacketHeader struct {
	Tssec uint32
	Tsusec uint32
	Caplen uint32
	Origlen uint32
}
