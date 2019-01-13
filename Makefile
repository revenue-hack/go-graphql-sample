.PHONY: setup run

all: main

main: setup

setup:
	docker-compose up -d
	go build -o bin/exec

run: setup
	./bin/exec

