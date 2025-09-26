.DEFAULT_GOAL := build

.PHONY:fmt vet build clean run
fmt: 
	go fmt ./...
vet: fmt
	go vet ./...
build: vet fmt run
	go build
clean:
	go clean
run: vet fmt
	go run .