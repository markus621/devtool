SHELL=/bin/bash

install:
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o $(GOPATH)/bin/devtool ./cmd/devtool/

unpack_demo: install
	rm -rf ./../demo
	cd ./../ && mkdir demo && cd demo && devtool dewep new github.com/1-2-3-4-5/6-7-8-9-0

pack_demo:
	rm -rf ./../demo/.idea
	rm -rf ./../demo/.vscode
	cd ./app/commands/dewep && static ./../../../../demo dewepNewProject
