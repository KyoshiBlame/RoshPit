FROM golang:1.26.9-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server ./cmd/server/main.go

FROM alpine:3.23

WORKDIR /app

COPY --from=builder /app/server .

CMD["./server"]