tests:
	go test -v -cover ./...

run:
	go run main.go

mock-install:
	go install github.com/golang/mock/mockgen@v1.6.0

build:
	go build -o ./bin/volume-take-home-assignment

docker-build:
	docker build -f ./docker/Dockerfile . -t volume-take-home-assignment

docker-run:
	docker run -p 8080:8080 -e PORT=8080 volume-take-home-assignment
