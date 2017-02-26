.PHONY: all
all: linux windows osx

.PHONY: linux
linux: *.go
	GOOS=linux GOARCH=amd64 go build -o binaries/microbadger_linux_64bit
	GOOS=linux GOARCH=386 go build -o binaries/microbadger_linux_32bit
	strip binaries/microbadger_linux_*

.PHONY: windows
windows: *.go
	goversioninfo -icon=logos/microbadger.ico
	GOOS=windows GOARCH=386  go build -ldflags -H=windowsgui -o binaries/microbadger_windows_32bit.exe
	GOOS=windows GOARCH=amd64  go build -ldflags -H=windowsgui -o binaries/microbadger_windows_64bit.exe
	rm resource.syso

.PHONY: osx
osx: *.go
	GOOS=darwin GOARCH=amd64 go build -o binaries/microbadger_osx_64bit
	GOOS=darwin GOARCH=386 go build -o binaries/microbadger_osx_32bit


.PHONY: clean
clean:
	rm -rf binaries/*
