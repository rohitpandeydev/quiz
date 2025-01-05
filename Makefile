.PHONY: all fmt vet build run

all: fmt vet build run

fmt:
	go fmt main.go db.go

vet: fmt
	go vet main.go db.go

build: vet
	go build -o bin/main main.go db.go

run: build
	./bin/main
