# Build stage
FROM golang:1.21-alpine as builder

WORKDIR /app
COPY . .

# Download dependencies
RUN go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o webapp .

# Runtime stage
FROM alpine:3.18

WORKDIR /app

# Copy binary and assets from builder
COPY --from=builder /app/webapp .
COPY --from=builder /app/static ./static

# Security best practices
RUN addgroup -S appgroup && \
    adduser -S appuser -G appgroup && \
    chown -R appuser:appgroup /app

USER appuser

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s \
  CMD wget --quiet --tries=1 --spider http://localhost:8080/ || exit 1

CMD ["./webapp"]