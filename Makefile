.PHONY: setup run

all: main

main: setup

setup:
	go build -o bin/exec

run: setup
	./bin/exec

