# Runtime stage
FROM alpine:3.18
WORKDIR /app

COPY --from=builder /app/webapp .
COPY --from=builder /app/static ./static

EXPOSE 8080
RUN ls -la /app/static/ && \
    echo "Проверка статических файлов:" && \
    [ -f /app/static/image.jpg ] && echo "image.jpg найден" || echo "image.jpg отсутствует" && \
    [ -f /app/static/image2.jpg ] && echo "image2.jpg найден" || echo "image2.jpg отсутствует" && \
    [ -f /app/static/styles.css ] && echo "styles.css найден" || echo "styles.css отсутствует"
CMD ["./webapp"]