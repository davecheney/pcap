#define _LARGEFILE_SOURCE
#define _LARGEFILE64_SOURCE
#define _FILE_OFFSET_BITS 64
#define _GNU_SOURCE

#include <sys/time.h>
#include <sys/times.h>
#include <sys/timex.h>
#include <pcap/pcap.h>

typedef struct pcap_file_header $FileHeader;
typedef struct pcap_pkthdr $PacketHeader;
