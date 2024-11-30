# Binary name
BINARY=repoview

# Installation paths
PREFIX ?= ${HOME}/.local/bin

# Go related variables
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

# Build target
.PHONY: build
build:
	@echo "Building..."
	@go build -o $(GOBIN)/$(BINARY)

# Clean target
.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -rf $(GOBIN)

# Install target
.PHONY: install
install: build
	@echo "Installing to $(PREFIX)..."
	@mkdir -p $(PREFIX)
	@cp $(GOBIN)/$(BINARY) $(PREFIX)/$(BINARY)

# Uninstall target
.PHONY: uninstall
uninstall:
	@echo "Uninstalling from $(PREFIX)..."
	@rm -f $(PREFIX)/$(BINARY)

# Test target
.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...

# Format target
.PHONY: format
format:
	@echo "Formatting code..."
	@go fmt ./...
