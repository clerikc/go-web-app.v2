# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Копируем сначала только файлы модуля
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем весь остальной код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o webapp .

# Runtime stage
FROM alpine:3.18
WORKDIR /app

# Копируем бинарник и статические файлы
COPY --from=builder /app/webapp .
COPY --from=builder /app/static ./static

EXPOSE 8080
CMD ["./webapp"]