FROM alpine:latest

RUN mkdir /app

COPY online-store /app

CMD ["/app/online-store"]
