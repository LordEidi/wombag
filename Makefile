# Makefile

#export GOPATH := $(shell pwd)

build:
	go build -o bin/wombagd cmd/wombagd/wombagd.go

	go build -o bin/wombagcli cmd/wombagcli/wombagcli.go


tidy:
	go mod tidy


run:
	go run cmd/wombagd/wombagd.go

compile:
#	echo "Compiling for every OS and Platform"
#	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
#	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
#	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go

.PHONY: help
all: help
help: Makefile
	@echo "help"