# syntax=docker/dockerfile:1

FROM golang:1.14 AS lang

WORKDIR /app

COPY . .

RUN mkdir executable
RUN go get -d ./... && go build -o executable ./...

EXPOSE 8080
ENTRYPOINT ["./executable/api"]
