BASE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
RELEASE_DIR := release
BINARY=go-tracing
all:
	$(MAKE) -C $(dir $(BASE_DIR)) build

.PHONY: clean
clean: 
	rm -rf ${RELEASE_DIR}

.PHONY: init
init: clean
	mkdir -p ${RELEASE_DIR}

.PHONY: build
build: init
	go build -o ${RELEASE_DIR}/${BINARY} -race -ldflags \
	'-X main.version=${VERSION} -X main.date=${DATE}' \
	./pkg
	@chmod +x ${RELEASE_DIR}/${BINARY}