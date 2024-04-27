# Define Go command and flags
GO = go
GOFLAGS = -ldflags="-s -w"

# Define the target executable
TARGET = twc

# Default target: build the executable
all: build

# Rule to build the target executable
build:
	mkdir -p dist
	$(GO) build $(GOFLAGS) -o dist/$(TARGET) *.go

# Clean target: remove the target executable
clean:
	rm -f ./dist

# Run target: build and run the target executable
run: build
	./dist/$(TARGET)

# Test target: run Go tests for the project
test:
	$(GO) test ./...
