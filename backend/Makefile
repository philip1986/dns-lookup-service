BINARY_NAME=dns-lookup-service

test:
	go test ./...

build:
	go build -o ${BINARY_NAME}

clean:
	go clean

test_and_build: test build
