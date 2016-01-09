BIN_DIR=bin

SERVER_BINARY=ghost
SERVER_BINARY_OUTPUT=${BIN_DIR}/${SERVER_BINARY}

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BENCH_SOURCES=./benchmark/ghostbench.go
BENCH_BINARY_OUTPUT=${BIN_DIR}/ghost-benchmark


ghost: $(SOURCES)
	go build -o ${SERVER_BINARY_OUTPUT} ghost_server.go

ghostbench: $(SOURCES)
	go build -o ${BENCH_BINARY_OUTPUT} ${BENCH_SOURCES}

.PHONY: install
install:
	go install ./...

.PHONY: clean
clean:
	if [ -f ${SERVER_BINARY_OUTPUT} ] ; then rm ${SERVER_BINARY_OUTPUT} ; fi
	if [ -f ${BENCH_BINARY_OUTPUT} ] ; then rm ${BENCH_BINARY_OUTPUT} ; fi

.PHONY: runbenchmark
runbenchmark:
	go test ./ghost -bench=. -benchmem
