BIN_DIR=bin

SERVER_BINARY=ghost
SERVER_BINARY_OUTPUT=${BIN_DIR}/${SERVER_BINARY}

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

ghost: $(SOURCES)
	go build -o ${SERVER_BINARY_OUTPUT} ghost_server.go

.PHONY: install
install:
	go install ./...

.PHONY: clean
clean:
	if [ -f ${SERVER_BINARY_OUTPUT} ] ; then rm ${SERVER_BINARY_OUTPUT} ; fi
