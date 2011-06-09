include $(GOROOT)/src/Make.inc

TARG=github.com/davecheney/pcap
GOFILES=\
	ethernet.go\
	mac.go\
	pcap.go\

GOFILES_darwin=\
	bpf.go\

GOFILES+=$(GOFILES_$(GOOS))

include $(GOROOT)/src/Make.pkg
