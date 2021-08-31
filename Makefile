#Docker shortcuts
.PHONY:br
br: buildd rund

.PHONY:buildd
buildd: build
	docker build -t rummy-go .

.PHONY:rund
rund:
	docker run rummy-go:latest

#Go shortcuts
.PHONY:all
all: fmt lint build

.PHONY:build
build:
	go build .

.PHONY:fmt
fmt:
	gofmt -s -w .

.PHONY:lint
lint:
	golangci-lint run ./...
