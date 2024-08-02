FROM golang:1.22.4-alpine as builder

WORKDIR /cmd

COPY ./ ./

RUN go mod tidy

WORKDIR /cmd/main

RUN go build main.go

CMD ["./main"]


