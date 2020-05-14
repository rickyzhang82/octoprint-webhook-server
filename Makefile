.PHONY: all go clean

all: go

go:
	go build -o octoprint-webhook-server.o launcher/*.go

install: go
	cp octoprint-webhook-server.o /usr/bin/
	cp resources/octoprint-webhook-server.service /usr/lib/systemd/system/

clean:
	rm octoprint-webhook-server.o

uninstall: clean
	rm /usr/bin/octoprint-webhook-server.o
	rm /usr/lib/systemd/system/octoprint-webhook-server.service