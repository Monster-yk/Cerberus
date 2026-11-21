APP_NAME := cerberus
BUILD_DIR := bin

.PHONY: all build run test fmt vet tidy clean help

all: build

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) ./

run:
	go run ./

test:
	go test ./...

fmt:
	gofmt -w *.go

vet:
	go vet ./...

tidy:
	go mod tidy

clean:
	rm -rf $(BUILD_DIR)

help:
	@echo "Available targets:"
	@echo "  make build  - Build the Cerberus binary"
	@echo "  make run    - Run Cerberus"
	@echo "  make test   - Run all Go tests"
	@echo "  make fmt    - Format Go source files"
	@echo "  make vet    - Run Go static analysis"
	@echo "  make tidy   - Clean Go module dependencies"
	@echo "  make clean  - Remove build artifacts"
