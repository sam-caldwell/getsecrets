.PHONY:all
.PHONY:build
.PHONY:clean
.PHONY:test

all: clean build test
	echo 'done'

clean:
	@rm -rf ./build
	@mkdir ./build
	@echo 'clean: done'

build:
	@echo 'building...'
	go build -o ./build/getsecrets main.go
	./build/getsecrets -version
	@echo 'build: done'


test:
	go test -v ./...
	go run main.go -version

version:
	./build/getsecrets -version