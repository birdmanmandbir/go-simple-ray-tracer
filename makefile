build:
	go build -o bin/simple-rt main.go

run: build
	./bin/simple-rt

test:
	go test ./... -count=1 -v

fmt:
	go fmt ./...