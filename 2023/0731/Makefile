EXPORT_NAME=hub.bacnet.tar.gz
name ?= test
toolname ?= bacnet-tools
os ?= linux
arch ?= amd64

version=$(shell git describe --abbrev=0 --tags)

clean:
	rm $(name) 2> /dev/null || true
	rm $(EXPORT_NAME) 2> /dev/null || true

build-pi:
	make build os=linux arch=arm

build-ubuntu:
	make build os=linux arch=amd64

build-windows:
	make build os=windows arch=amd64 name

build:
	GOOS=$(os) GOARCH=$(arch) go build -ldflags="-X main.version=$(version)" -o $(name) .

build-tool-pi:
	make build-tool os=linux arch=arm

build-tool-ubuntu:
	make build-tool os=linux arch=amd64

build-tool-windows:
	make build-tool os=windows arch=amd64 name

build-tool:
	GOOS=$(os) GOARCH=$(arch) go build -ldflags="-X main.version=$(version)" -o $(toolname) ./tools
