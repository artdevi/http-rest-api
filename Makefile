.PHONY: build
build:
	go build -v ./cmd/apiserver
.PHONY: test
test:
	go test -v -timeout 30s ./...

.PHONY: run
run: 
	go build -v ./cmd/apiserver
	./apiserver.exe

.DEFAULT_GOAL := build