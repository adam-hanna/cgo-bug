all: build
              
.PHONY: build
build:
	@go build -o ./lib/lib.so -buildmode=c-shared ./lib/lib.go
	
.PHONY: run
run:
	@go run main.go
