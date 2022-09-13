BINARY_NAME=pkb
DIR=./...

.PHONY: build
build:
	@go build -o ${BINARY_NAME} .

.PHONY: fmt
fmt:
	@go fmt ${DIR}

.PHONY: lint
lint:
	@golangci-lint run -v ./...

.PHONY: test
test:
	@CGO_ENABLED=1 go test ${DIR} -race -cover
