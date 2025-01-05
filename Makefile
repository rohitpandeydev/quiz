.PHONY: all fmt vet build run

all: fmt vet build run

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o bin/main cmd/main/main.go

run: build
	./bin/main
