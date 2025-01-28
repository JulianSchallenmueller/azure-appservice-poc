FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY ./main.go .
COPY ./go.mod .
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o go-poc

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/go-poc .
CMD ["./go-poc"]

EXPOSE 80:80/tcp

