.PHONY: deploy build clean config run

all:
	make setup
	make build
	make exportvariables
	make run

setup:
	go mod vendor

build:
	go build -v -o main ./cmd/app/main.go

exportvariables:
	export CGO_ENABLED=1

run:
	./main --configpath ./cmd/app/config.yaml

