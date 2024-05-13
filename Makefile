.SILENT:
.DEFAULT_GOAL := run-fast

.PHONY: build
build:
	go build -o ./build/chess ./cmd/chess

.PHONY: run
run:
	./build/chess

.PHONY: run-fast
run-fast:
	go run ./cmd/chess
