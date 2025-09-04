BINARY_NAME=synapse-ai

# Go build flags (optional: -ldflags="-s -w" strips debug info for smaller binaries)
BUILD_FLAGS=

# Default target: build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	go build $(BUILD_FLAGS) -o bin/$(BINARY_NAME) ./cmd/server

# Run the project directly (without manually building first)
run:
	@echo "Running $(BINARY_NAME)..."
	go run ./cmd/server

# Run tests
runTest:
	@echo "Running tests..."
	go test ./... -v

# Clean up build artifacts
clean:
	@echo "Cleaning..."
	rm -rf bin

# Install dependencies (tidy up go.mod / go.sum)
deps:
	@echo "Tidying up modules..."
	go mod tidy

# Rebuild from scratch
rebuild: clean build
