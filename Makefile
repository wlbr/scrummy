LINKERFLAGS = -X main.Version=`git describe --tags --always --long --dirty` -X main.Buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S_UTC'`
TARGET = scrummy

all: clean build

.PHONY: clean
clean:
	rm -rf bin/
	rm -f $(TARGET)
	rm -f gotils/loglevel_string.go

generate:
	go generate gotils/loglevel.go

build: generate
	go build -ldflags "$(LINKERFLAGS)" cmd/$(TARGET)/$(TARGET).go

run: generate
	#go run -ldflags "$(LINKERFLAGS)" cmd/$(TARGET)/$(TARGET).go 
	go run cmd/$(TARGET)/$(TARGET).go -LogLevel Warn

debug: generate
	dlv debug cmd/$(TARGET)/$(TARGET).go -RunAsServer

buildcmd: clean generate
	mkdir -p bin
	GOOS=linux GOARCH=amd64   go build -o bin/$(TARGET)_linux_x86_64 -ldflags "$(LINKERFLAGS)" cmd/$(TARGET)/$(TARGET).go
	GOOS=darwin GOARCH=amd64  go build -o bin/$(TARGET)_osx_x86_64   -ldflags "$(LINKERFLAGS)" cmd/$(TARGET)/$(TARGET).go
	GOOS=windows GOARCH=386   go build -o bin/$(TARGET)_win32        -ldflags "$(LINKERFLAGS)" cmd/$(TARGET)/$(TARGET).go
	GOOS=windows GOARCH=amd64 go build -o bin/$(TARGET)_win64        -ldflags "$(LINKERFLAGS)" cmd/$(TARGET)/$(TARGET).go



deploy: rbuild
	rsync -a --progress $(TARGET) $(HOST):./fcgi-bin/

dep:
	go get -u github.com/360EntSecGroup-Skylar/excelize
	go get -u github.com/wcharczuk/go-chart
	#go get -u github.com/wlbr/enumer
	go get -u github.com/alvaroloes/enumer
