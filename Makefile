# Binary name
BINARY=git-track

# Installation paths
PREFIX ?= ${HOME}/.local
INSTALL_PATH = $(PREFIX)/bin

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
	@echo "Installing to $(INSTALL_PATH)..."
	@mkdir -p $(INSTALL_PATH)
	@cp $(GOBIN)/$(BINARY) $(INSTALL_PATH)/$(BINARY)

# Uninstall target
.PHONY: uninstall
uninstall:
	@echo "Uninstalling from $(INSTALL_PATH)..."
	@rm -f $(INSTALL_PATH)/$(BINARY)

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
