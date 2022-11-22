tests:
	go test -v -cover ./...

run:
	go run main.go

build:
	go build -o ./bin/volume-take-home-assignment

docker-build:
	docker build -f ./docker/Dockerfile . -t volume-take-home-assignment

docker-run:
	docker run -d -p 8080:8080 -e PORT=8080 volume-take-home-assignment
