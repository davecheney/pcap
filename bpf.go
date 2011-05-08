package pcap

import (
	"os"
	"syscall"
	"unsafe"
)

const (
	device = "/dev/bpf0"
)

type reader struct {
	fd     int
	buflen int // buffer size supplied by bpf
}

type Capture struct {
	header  syscall.BpfHdr
	payload []byte
}

func (r *reader) ReadPacket() (*Capture, os.Error) {
	buf := make([]byte, r.buflen)
	n, e := syscall.Read(r.fd, buf)
	if e != 0 {
		return nil, &os.PathError{"read", device, os.Errno(e)}
	}
	buf = buf[:n]
	header := *(*syscall.BpfHdr)(unsafe.Pointer(&buf[0]))
	capture := &Capture{
		header:  header,
		payload: buf[header.Hdrlen : uint32(header.Hdrlen)+header.Caplen],
	}
	return capture, nil
}

func (r *reader) Close() os.Error {
	syscall.Close(r.fd)
	return nil // TODO(dfc)
}


func Open() (PacketReader, os.Error) {
	fd, e := syscall.Open(device, os.O_RDONLY|syscall.O_CLOEXEC, 0666)
	if e != 0 {
		return nil, &os.PathError{"open", device, os.Errno(e)}
	}
	var data [16]byte
	data[0] = 'e'
	data[1] = 'n'
	data[2] = '1'

	var len uint32
	var immediate uint32 = 1
	var promisc uint32 = 1
	if err := os.NewSyscallError("ioctl", syscall.Ioctl(fd, syscall.BIOCGBLEN, uintptr(unsafe.Pointer(&len)))); err != nil {
		return nil, err
	}
	if err := os.NewSyscallError("ioctl", syscall.Ioctl(fd, syscall.BIOCSBLEN, uintptr(unsafe.Pointer(&len)))); err != nil {
		return nil, err
	}
	if err := os.NewSyscallError("ioctl", syscall.Ioctl(fd, syscall.BIOCIMMEDIATE, uintptr(unsafe.Pointer(&immediate)))); err != nil {
		return nil, err
	}
	if err := os.NewSyscallError("ioctl", syscall.Ioctl(fd, syscall.BIOCSETIF, uintptr(unsafe.Pointer(&data[0])))); err != nil {
		return nil, err
	}
	if err := os.NewSyscallError("ioctl", syscall.Ioctl(fd, syscall.BIOCPROMISC, uintptr(unsafe.Pointer(&promisc)))); err != nil {
		return nil, err
	}
	return &reader{fd, int(len)}, nil
}
