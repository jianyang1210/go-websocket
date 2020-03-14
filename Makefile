.PHONY: all fmt dep lint test build apidoc

export GO111MODULE=on
export GOPROXY=https://goproxy.cn

ExecutableFile=gowebsocket

all: clean fmt dep lint test build

fmt:
	@go fmt ./...

dep:
	@go mod download && go mod tidy

lint:
	@golangci-lint run

test:
	@go test -v -cover ./...

build:
	@echo ${VER_FLAGS}
	@GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o ${ExecutableFile} main.go

clean:
	@rm -rf bin ${ExecutableFile}
