include $(GOROOT)/src/Make.inc

TARG=github.com/davecheney/pcap
GOFILES=\
	pcap.go\

GOFILES_darwin=\
	bpf.go\

GOFILES+=$(GOFILES_$(GOOS))

include $(GOROOT)/src/Make.pkg
