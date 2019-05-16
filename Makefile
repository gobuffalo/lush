TAGS ?= "sqlite"
GO_BIN ?= go

gen:
	pigeon -o internal/parser/parse.go internal/parser/lush.peg
	go install -v ./cmd/lush

install: gen
	$(GO_BIN) install -tags ${TAGS} -v ./cmd/lush
	make tidy

tidy:
ifeq ($(GO111MODULE),on)
	$(GO_BIN) mod tidy
else
	echo skipping go mod tidy
endif

deps:
	go get github.com/mna/pigeon
	$(GO_BIN) get -tags ${TAGS} -t ./...
	make tidy

build: gen
	$(GO_BIN) build -v .
	make tidy

test: gen
	$(GO_BIN) test -cover -tags ${TAGS} ./...
	make tidy

ci-deps:
	$(GO_BIN) get -tags ${TAGS} -t ./...

ci-test:
	$(GO_BIN) test -tags ${TAGS} -race ./...

update:
	$(GO_BIN) get -u -tags ${TAGS}
	make tidy
	make test
	make install
	make tidy

release-test:
	$(GO_BIN) test -tags ${TAGS} -race ./...
	make tidy

release:
	$(GO_BIN) get github.com/gobuffalo/release
	make tidy
	release -y -f version.go
	make tidy

examples: install
	lush ./examples/big.lush ./examples/returns.lush ./examples/errors.lush
