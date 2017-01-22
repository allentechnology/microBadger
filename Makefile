.PHONY: all
all: linux windows

.PHONY: linux
linux: 
	GOOS=linux GOARCH=amd64 go build -o binaries/microbadger_linux_64bit
	GOOS=linux GOARCH=386 go build -o binaries/microbadger_linux_32bit
	strip binaries/microbadger_linux_*

.PHONY: windows
windows:
	GOOS=windows GOARCH=amd64 go build -o binaries/microbadger_windows_64bit.exe
	GOOS=windows GOARCH=386 go build -o binaries/microbadger_windows_32bit.exe

.PHONY: clean
clean:
	rm -rf binaries/*
