.PHONY: dev build generate install image release profile bench test clean setup

CGO_ENABLED=0
VERSION=$(cat version)
BIN=bin
SGO=Server.go
CGO=Client.go

SBIN=bin/Server
CBIN=bin/Client

all: dev

dev: build

build: generate
	go build -o $(CBIN) $(CGO)
	go build -o $(SBIN) $(SGO)

bench: build
	sh bench/BenchMark.sh
