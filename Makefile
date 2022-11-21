
unit:
	go test -v -tags=unit -cover ./...

e2e:
	go test -v -tags=integration -cover ./...

tests:
	go test -v -tags=integration,unit -cover ./...

run:
	go run main.go

build:
	go build -o ./bin/volume-take-home-assignment