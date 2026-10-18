# Stage 1: Builder
FROM golang:1.26-bookworm AS builder

WORKDIR /workspace

# Сначала копируем только go.mod и go.sum для кэширования слоёв Docker
COPY proto/go.mod proto/go.sum* ./proto/
COPY event-api/go.mod event-api/go.sum ./event-api/

# Качаем зависимости
WORKDIR /workspace/event-api
RUN go mod edit -replace github.com/your-username/roshpit/proto=../proto && go mod download

# Копируем исходники proto и самого сервиса
WORKDIR /workspace
COPY proto/ ./proto/
COPY event-api/ ./event-api/

# Собираем статический бинарник
WORKDIR /workspace/event-api
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /bin/event-api ./cmd/server/main.go

# Stage 2: Runner (минимальный образ)
FROM alpine:3.24

RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app

COPY --from=builder /bin/event-api /app/event-api
# Копируем миграции, если они запускаются из кода при старте
COPY --from=builder /workspace/event-api/migrations /app/migrations

EXPOSE 8080
CMD ["/app/event-api"]