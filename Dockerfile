FROM golang:1.21 AS builder
ENV CGO_ENABLED 0
ADD . /app
WORKDIR /app
RUN go build -ldflags "-s -w" -v -o github-status-exporter .

FROM alpine:3
RUN apk update && \
    apk add openssl tzdata && \
    rm -rf /var/cache/apk/* \
    && mkdir /app

WORKDIR /app

COPY --from=builder /app/github-status-exporter /app/github-status-exporter

RUN chown -R nobody /app \
    && chmod 500 /app/github-status-exporter \
    && chmod -R 700 /app/data

USER nobody
ENTRYPOINT ["/app/github-status-exporter"]