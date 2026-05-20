APP_NAME=shipyard-api

run:
	go run cmd/api/main.go

build:
	go build -o bin/$(APP_NAME) cmd/api/main.go

test:
	go test ./...

fmt:
	go fmt ./...

clean:
	rm -rf bin