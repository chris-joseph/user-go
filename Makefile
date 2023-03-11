VERSION  = `git describe --tags`
BUILD    = `date +%FT%T%z`
LDFLAGS  = -ldflags "-X main.AppVersion=${VERSION} -X main.BuildTime=${BUILD}"

.PHONY: init
init:
	npm i
	npm run prepare
	brew list golangci-lint || brew install golangci-lint
	brew list go-swagger || brew tap go-swagger/go-swagger && brew install go-swagger
	go install github.com/golang/mock/mockgen@v1.6.0
	npm i -g redoc-cli

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

.PHONY: gendocs
gendocs:
	swagger generate spec -i docs/swagger.base.yaml -o ./docs/__swagger.yaml --scan-models
	redoc-cli build docs/__swagger.yaml -o docs/__index.html
