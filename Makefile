build:
	@go build -o bin/coincrowd

run: build
	@./bin/coincrowd

test:
	@go test -v ./...