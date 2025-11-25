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
	@go build -ldflags="-w -s" -o updep ./cmd/updep

.PHONY: start
start:
	@./tmp/updep

.PHONY: clean
clean:
	@rm -rf ./updep

