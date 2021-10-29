FROM golang:latest

WORKDIR /home/go-rest-api

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go", "run", "server.go"]