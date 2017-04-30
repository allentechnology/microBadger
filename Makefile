.PHONY: all
all: linux windows osx

.PHONY: windows-dependencies
windows-dependencies:
	go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo

.PHONY: embed-assets
embed-assets:
	go get github.com/jteeuwen/go-bindata/...
	go-bindata ./logos/microBadger_headert.png ./webpage.html

.PHONY: linux
linux: *.go embed-assets
	GOOS=linux GOARCH=amd64 go build -o binaries/microbadger_linux_64bit
	GOOS=linux GOARCH=386 go build -o binaries/microbadger_linux_32bit
	strip binaries/microbadger_linux_*

.PHONY: windows 
windows: *.go windows-dependencies embed-assets
	goversioninfo -icon=logos/microbadger.ico
	GOOS=windows GOARCH=386  go build -ldflags -H=windowsgui -o binaries/microbadger_windows_32bit.exe
	GOOS=windows GOARCH=amd64  go build -ldflags -H=windowsgui -o binaries/microbadger_windows_64bit.exe
	rm resource.syso

.PHONY: osx
osx: *.go embed-assets
	GOOS=darwin GOARCH=amd64 go build -o binaries/microbadger_osx_64bit
	GOOS=darwin GOARCH=386 go build -o binaries/microbadger_osx_32bit


.PHONY: clean
clean:
	rm -rf binaries/*
