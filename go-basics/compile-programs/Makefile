run:
	@go run main.go

build-macos-amd64:
	@GOOS=darwin GOARCH=amd64 go build -o sum-macos-amd64 main.go

build-macos-arm64:
	@GOOS=darwin GOARCH=arm64 go build -o sum-macos-arm64 main.go

build-linux-amd64:
	@GOOS=linux GOARCH=amd64 go build -o sum-linux-amd64 main.go

build-windows-amd64:
	@GOOS=windows GOARCH=amd64 go build -o sum-windows-amd64 main.go

help:
	@echo "Usage: make <target>"
	@echo "Targets:"
	@echo "  run - Run the program"
	@echo "  build-macos-amd64 - Build for macOS (amd64)"
	@echo "  build-macos-arm64 - Build for macOS (arm64)"
	@echo "  build-linux-amd64 - Build for Linux (amd64)"
