BIN_DIR=bin
CMD_DIR=cmd

SERVER_BINARY=ghost-server
SERVER_DIR=$(CMD_DIR)/${SERVER_BINARY}
SERVER_BINARY_OUTPUT=${BIN_DIR}/${SERVER_BINARY}

CLI_BINARY=ghost-cli
CLI_DIR=$(CMD_DIR)/${CLI_BINARY}
CLI_BINARY_OUTPUT=${BIN_DIR}/${CLI_BINARY}

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BENCH_BINARY=ghost-benchmark
BENCH_DIR=$(CMD_DIR)/${BENCH_BINARY}
BENCH_BINARY_OUTPUT=${BIN_DIR}/${BENCH_BINARY}


ghost: $(SOURCES)
	go build -o ${SERVER_BINARY_OUTPUT} ${SERVER_DIR}/server.go

ghost-cli: $(SOURCES)
	go build -o ${CLI_BINARY_OUTPUT} ${CLI_DIR}/cli.go

ghost-benchmark: $(SOURCES)
	go build -o ${BENCH_BINARY_OUTPUT} ${BENCH_DIR}/ghostbench.go

all: ghost ghost-benchmark ghost-cli

.PHONY: install
install:
	go install ./cmd/ghost-benchmark
	go install ./cmd/ghost-cli
	go install ./cmd/ghost-server

.PHONY: clean
clean:
	if [ -f ${SERVER_BINARY_OUTPUT} ] ; then rm ${SERVER_BINARY_OUTPUT} ; fi
	if [ -f ${BENCH_BINARY_OUTPUT} ] ; then rm ${BENCH_BINARY_OUTPUT} ; fi
	if [ -f ${CLI_BINARY_OUTPUT} ] ; then rm ${CLI_BINARY_OUTPUT} ; fi
.PHONY: runbenchmark
runbenchmark:
	go test ./ghost -bench=. -benchmem
