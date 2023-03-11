VERSION  = `git describe --tags`
BUILD    = `date +%FT%T%z`
LDFLAGS  = -ldflags "-X main.AppVersion=${VERSION} -X main.BuildTime=${BUILD}"

.PHONY: init
init:
	npm i
	npm run prepare
	brew install golangci-lint
	go install github.com/golang/mock/mockgen@v1.6.0

.PHONY: all
all: lint test build

.PHONY: run
run:
	go run ${LDFLAGS} cmd/api/main.go

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -v -cover -coverprofile=cover.out -test.bench=. -benchmem ./...

.PHONY: build
build:
	go build ${LDFLAGS} -o bin/go-template-be-api cmd/api/main.go

.PHONY: testcoverage
testcoverage:
	./scripts/test_coverage_threshold.sh
