.SILENT:
.DEFAULT_GOAL := fast-run

.PHONY: build
build:
	go build -o ./build/chess ./cmd/chess

.PHONY: run
run:
	./build/chess

.PHONY: fast-run
fast-run:
	go run ./cmd/chess
