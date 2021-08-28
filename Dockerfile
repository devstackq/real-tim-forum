FROM golang:latest
LABEL maintainer="devstackq"
RUN mkdir /build
COPY go.mod go.sum ./
RUN go mod download 
COPY . .
RUN CGO_ENABLED =0 GOOS=linux go build -o installsuffix cgo -main ./cmd
FROM alpine:latest
RUN apk --no-cache add ca-certificates
EXPOSE 6969
CMD ["./main"]
WORKDIR /build