include $(GOROOT)/src/Make.inc

TARG= htdigest
GOFILES= main.go htfile.go

include $(GOROOT)/src/Make.cmd
