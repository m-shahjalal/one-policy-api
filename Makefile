run:
	go run ./cmd/main.go

build:
	go build -o bin/app ./cmd/main.go

test:
	go test ./...

dev:
	~/go/bin/air 