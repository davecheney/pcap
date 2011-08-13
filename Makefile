include $(GOROOT)/src/Make.inc

TARG=github.com/davecheney/pcap
GOFILES=\
	ethernet.go\
	mac.go\
	pcap.go\

GOFILES_darwin=\
	bpf.go\

GOFILES_amd64=\
	ztypes_amd64.go\

GOFILES_386=\
	ztypes_386.go\

GOFILES+=$(GOFILES_$(GOOS))

GOFILES+=$(GOFILES_$(GOARCH))

CLEAN_FILES+=ztypes_*.go

include $(GOROOT)/src/Make.pkg

ztypes_386.go: types.c
	godefs -gpcap -f -m32 $^ | gofmt > $@

ztypes_amd64.go: types.c
	godefs -gpcap -f -m64 $^ | gofmt > $@

