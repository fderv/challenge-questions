FROM golang:1.14.7

ENV GO111MODULE=on

WORKDIR /server
ADD server.go go.mod go.sum /server/
ADD user /server/user

RUN go mod download

EXPOSE 50051

CMD ["go", "run", "/server/server.go"]