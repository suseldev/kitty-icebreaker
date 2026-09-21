.PHONY: all build kittyice kittyinit clean

BUILD_DIR := build

all: build

build: kittyice kittyinit

kittyice:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/kittyice ./cmd/kittyice

kittyinit:
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o $(BUILD_DIR)/kittyinit ./cmd/kittyinit

clean:
	rm -rf $(BUILD_DIR)

