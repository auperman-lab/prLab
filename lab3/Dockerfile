FROM golang:1.23.2-alpine3.20
WORKDIR /app

COPY manager ./manager

COPY go.mod go.sum main.go ./

RUN go mod download && \
    go build -o main

EXPOSE 8080


CMD ["./main"]