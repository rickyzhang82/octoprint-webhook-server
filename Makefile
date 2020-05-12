.PHONY: all go clean

all: go

go:
	go build -o app launcher/*.go

clean:
	rm app

