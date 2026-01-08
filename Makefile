.PHONY: build install clean test

# Build the binary
build:
	go build -o bii .

# Build for multiple platforms
build-all:
	GOOS=linux GOARCH=amd64 go build -o dist/bii-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -o dist/bii-linux-arm64 .
	GOOS=darwin GOARCH=amd64 go build -o dist/bii-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -o dist/bii-darwin-arm64 .

# Install to ~/.local/bin
install: build
	mkdir -p ~/.local/bin
	cp bii ~/.local/bin/
	@echo "✅ bii installed to ~/.local/bin/bii"
	@echo "Make sure ~/.local/bin is in your PATH"

# Clean build artifacts
clean:
	rm -f bii
	rm -rf dist/
	rm -rf test-data/

# Run tests
test:
	go test -v ./...

# Format code
fmt:
	go fmt ./...

# Run linter
lint:
	go vet ./...
