#!make
build_path = build
project_name = csv-parser
install_path := $(shell go env GOPATH)

run:
	go run main.go
.PHONY: run

build:
	go build -o $(build_path)/$(project_name)
.PHONY: build

clean:
	rm -rf $(build_path)
.PHONY: clean

install:
	go install
.PHONY: install

uninstall:
	rm -f $(install_path)/bin/$(project_name)
.PHONY: uninstall

fmt:
	go fmt
.PHONY: fmt