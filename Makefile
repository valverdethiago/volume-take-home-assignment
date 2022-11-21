tests:
	go test -v -cover ./...

run:
	go run main.go

build:
	go build -o ./bin/volume-take-home-assignment