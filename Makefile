BINARY_NAME=pkb
DIR=./...
VERSION ?= $(shell head -n 1 VERSION)

.PHONY: build
build:
	@go build -ldflags "-X github.com/thaffenden/pkb/cmd.Version=${VERSION}" -o ${BINARY_NAME} .

.PHONY: fmt
fmt:
	@go fmt ${DIR}

.PHONY: install
install: build
	@sudo cp ./${BINARY_NAME} /usr/bin/${BINARY_NAME}

.PHONY: lint
lint:
	@golangci-lint run -v ./...

.PHONY: push-tag
push-tag:
	@git tag -a ${VERSION} -m "Release ${VERSION}"
	@git push origin ${VERSION}

.PHONY: test
test:
	@CGO_ENABLED=1 go test ${DIR} -race -cover
