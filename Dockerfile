# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Копируем только файлы, необходимые для загрузки зависимостей
COPY go.mod ./
RUN go mod download

# Копируем остальные файлы и собираем
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o webapp .

# Runtime stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/webapp .
COPY --from=builder /app/static ./static

EXPOSE 8080
CMD ["./webapp"]