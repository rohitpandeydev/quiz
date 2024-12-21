.PHONY: all fmt vet build run

all: fmt vet build run

fmt:
	go fmt main.go

vet: fmt
	go vet main.go

build: vet
	go build -o bin/main main.go

run: build
	./bin/main
