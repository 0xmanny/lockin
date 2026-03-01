BINARY_NAME=lockin
CLI_DIR=cli
BUILD_DIR=build

VERSION ?= dev
LDFLAGS=-ldflags "-s -w -X main.version=$(VERSION)"

.PHONY: build build-all install clean

build:
	cd $(CLI_DIR) && go build $(LDFLAGS) -o ../$(BINARY_NAME) .

build-all: clean
	mkdir -p $(BUILD_DIR)
	cd $(CLI_DIR) && GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o ../$(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 .
	cd $(CLI_DIR) && GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o ../$(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 .

install: build
	sudo cp $(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)
	@echo "Installed $(BINARY_NAME) to /usr/local/bin/"

clean:
	rm -f $(BINARY_NAME)
	rm -rf $(BUILD_DIR)
	rm -rf ~/.lockin/

dev:
	$(MAKE) clean
	$(MAKE) build
	./$(BINARY_NAME)