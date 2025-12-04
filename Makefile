.PHONY: build install uninstall test clean help

APP_NAME=go-projo
VERSION=1.0.0

build:
	go build -o ${APP_NAME} .

install:
	go install

uninstall:
	rm -f ${GOPATH}/bin/${APP_NAME}

test:
	go test -v ./...

clean:
	rm -f ${APP_NAME}
	go clean

help:
	@echo "Available targets:"
	@echo "  build      - Build the application"
	@echo "  install    - Install to GOPATH/bin"
	@echo "  uninstall  - Remove from GOPATH/bin"
	@echo "  test       - Run tests"
	@echo "  clean      - Clean build artifacts"
	@echo "  help       - Show this help message"
