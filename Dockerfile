FROM golang:latest
RUN mkdir /build
ADD go.mod go.sum cmd/main.go /build/
WORKDIR /build
RUN go build