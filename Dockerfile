# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# 1. Сначала копируем ТОЛЬКО go.mod
COPY go.mod .

# 2. Скачиваем зависимости (работает без go.sum)
RUN go mod download

# 3. Копируем остальные файлы проекта
COPY . .

# 4. Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o webapp .

# Runtime stage
FROM alpine:3.18
WORKDIR /app

COPY --from=builder /app/webapp .
COPY --from=builder /app/static ./static

EXPOSE 8080
RUN ls -la /app/static/ && \
    echo "Проверка статических файлов:" && \
    [ -f /app/static/image.jpg ] && echo "image.jpg найдено" || echo "image.jpg отсутствует" && \
    [ -f /app/static/image2.jpg ] && echo "image2.jpg найдено" || echo "image2.jpg отсутствует" && \
    [ -f /app/static/styles.css ] && echo "styles.css найден" || echo "styles.css отсутствует"
CMD ["./webapp"]