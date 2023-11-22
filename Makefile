.PHONY: all windows linux macos-intel macos-arm

all: windows linux macos-intel macos-arm

windows:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o monitor_qubic-windows-amd64.exe

linux:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o monitor_qubic-linux-amd64

macos-intel:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o monitor_qubic-macos-amd64

macos-arm:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o monitor_qubic-macos-arm64