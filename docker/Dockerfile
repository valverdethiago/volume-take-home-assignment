FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /env
COPY env/app.env ./app.env

WORKDIR /

COPY . .

RUN go mod download

RUN go build -o /bin/volume-take-home-assignment

EXPOSE ${PORT}

CMD ["/bin/volume-take-home-assignment"]
