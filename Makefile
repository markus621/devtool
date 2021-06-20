SHELL=/bin/bash
TOOLS_BIN=$(shell pwd)/.tools

install:
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o $(GOPATH)/bin/devtool ./cmd/devtool/

unpack_demo: install
	rm -rf ./../demo
	cd ./../ && mkdir demo && cd demo && devtool dewep new github.com/1-2-3-4-5/6-7-8-9-0

pack_demo:
	rm -rf ./../demo/.idea
	rm -rf ./../demo/.vscode
	rm -rf ./../demo/.tools
	rm -rf ./../demo/*.out
	cd ./app/commands/dewep && static ./../../../../demo dewepNewProject

ci:
	go mod download
	go build -race -v -o /tmp/devtool ./cmd/devtool/
	rm -rf $(TOOLS_BIN)
	mkdir -p $(TOOLS_BIN)
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(TOOLS_BIN) v1.38.0
	$(TOOLS_BIN)/golangci-lint -v run ./...
	go test -race -v ./...
