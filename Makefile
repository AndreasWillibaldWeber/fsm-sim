GO ?= go
BUILD_FOLDER = "./build"
MAIN_FILE = "cmd/main.go"
OUT_FILE = "fsm"

all: build

build:
	if [ -d $(BUILD_FOLDER) ]; then rm -rf $(BUILD_FOLDER); fi
	mkdir $(BUILD_FOLDER)
	GOOS=linux GOARCH=amd64 $(GO) build $(GOFLAGS) -o $(BUILD_FOLDER)/${OUT_FILE} $(MAIN_FILE)

clean:
	if [ -d $(BUILD_FOLDER) ]; then rm -rf $(BUILD_FOLDER); fi

.PHONY: build clean
