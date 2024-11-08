PACKAGE_NAME := $(shell go list .)
PROGRAM_NAME := $(notdir $(PACKAGE_NAME))

CURRENT_VERSION := $(shell git describe --tags --abbrev=0)
CURRENT_REVISION := $(shell git rev-parse --short HEAD)
BUILD_LDFLAGS := " \
	-s -w \
	-X $(PACKAGE_NAME).version=$(CURRENT_VERSION) \
	-X $(PACKAGE_NAME).revision=$(CURRENT_REVISION) \
	"

.PHONY: test
test:
	go test ./..

.PHONY: build
build:
	go build -ldflags=$(BUILD_LDFLAGS) ./cmd/$(PROGRAM_NAME)

.PHONY: clean
clean:
	rm ./gopische
