BIN_DIR=bin

SERVER_BINARY=ghost
SERVER_BINARY_OUTPUT=${BIN_DIR}/${SERVER_BINARY}

CLI_BINARY=ghost-cli
CLI_BINARY_OUTPUT=${BIN_DIR}/${CLI_BINARY}

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BENCH_SOURCES=./ghost-benchmark/ghostbench.go
BENCH_BINARY_OUTPUT=${BIN_DIR}/ghost-benchmark


ghost: $(SOURCES)
	go build -o ${SERVER_BINARY_OUTPUT} ghost_server.go
	go build -o ${CLI_BINARY_OUTPUT} ghost_cli.go

ghost-benchmark: $(SOURCES)
	go build -o ${BENCH_BINARY_OUTPUT} ${BENCH_SOURCES}

all: ghost ghost-benchmark

.PHONY: install
install:
	go install ./...

.PHONY: clean
clean:
	if [ -f ${SERVER_BINARY_OUTPUT} ] ; then rm ${SERVER_BINARY_OUTPUT} ; fi
	if [ -f ${BENCH_BINARY_OUTPUT} ] ; then rm ${BENCH_BINARY_OUTPUT} ; fi
	if [ -f ${CLI_BINARY_OUTPUT} ] ; then rm ${CLI_BINARY_OUTPUT} ; fi
.PHONY: runbenchmark
runbenchmark:
	go test ./ghost -bench=. -benchmem
