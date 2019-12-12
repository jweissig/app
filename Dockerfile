FROM alpine:3.10.3
RUN mkdir -p /app
ADD web /app/web
WORKDIR /app
ENTRYPOINT ["/app/web"]
