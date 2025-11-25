.PHONY: all
all: build

.PHONY: format
format:
	@go fmt ./...

.PHONY: run
run:
	@go run ./cmd/updep

.PHONY: build
build:
	@go build -ldflags="-w -s" -o main ./cmd/updep

.PHONY: start
start:
	@./tmp/cli

.PHONY: clean
clean:
	@rm -rf ./tmp

