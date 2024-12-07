FROM golang:1.23.2-alpine3.20 as builder

WORKDIR /app

COPY raft/ ./raft/
COPY go.mod main.go ./

RUN go mod download && \
    go build -o main
EXPOSE 9001-9005

CMD ["./main"]
