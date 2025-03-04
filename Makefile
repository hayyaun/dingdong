# Define the binary name
BINARY_NAME = dingdong

# Go build flags for optimization
BUILD_FLAGS = -ldflags="-s -w"

# Default target: Build for the current OS
all: build

# Build the binary for the current OS
build:
	@echo "🔨 Building $(BINARY_NAME)..."
	go build $(BUILD_FLAGS) -o $(BINARY_NAME) main.go
	@echo "✅ Build complete: $(BINARY_NAME)"

# Cross-compile for different OS/Architectures
build-linux:
	GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -o $(BINARY_NAME)-linux main.go
build-windows:
	GOOS=windows GOARCH=amd64 go build $(BUILD_FLAGS) -o $(BINARY_NAME).exe main.go
build-mac:
	GOOS=darwin GOARCH=amd64 go build $(BUILD_FLAGS) -o $(BINARY_NAME)-mac main.go

# Install the binary to GOBIN (global install)
install:
	@echo "📦 Installing $(BINARY_NAME)..."
	go install
	@echo "✅ Installed to $(shell go env GOBIN)/$(BINARY_NAME)"

# Clean up built binaries
clean:
	@echo "🧹 Cleaning up..."
	rm -f $(BINARY_NAME) $(BINARY_NAME).exe $(BINARY_NAME)-linux $(BINARY_NAME)-mac
	@echo "✅ Clean complete"

# List all make commands
help:
	@echo "Available commands:"
	@echo "  make build        - Build for current OS"
	@echo "  make build-linux  - Build for Linux (amd64)"
	@echo "  make build-windows - Build for Windows (amd64)"
	@echo "  make build-mac    - Build for macOS (amd64)"
	@echo "  make install      - Install to GOBIN"
	@echo "  make clean        - Remove built binaries"
	@echo "  make help         - Show this help"
