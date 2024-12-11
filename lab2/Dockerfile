FROM golang:1.23.2-alpine3.20
WORKDIR /app

COPY internal ./internal
COPY pkg/ ./pkg/
COPY cmd/ ./cmd/
COPY raft/ ./raft/
COPY .env go.mod go.sum main.go ./

RUN go mod download && \
    go build -o main

EXPOSE 2003
EXPOSE 2001
EXPOSE 9001-9005


CMD ["./main"]