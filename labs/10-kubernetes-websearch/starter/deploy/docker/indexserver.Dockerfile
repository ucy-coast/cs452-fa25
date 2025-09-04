
FROM golang:1.24 AS builder
WORKDIR /app
COPY . .
RUN go build -o bin/indexserver ./cmd/indexserver

FROM ubuntu:22.04

# Install certificates and any other needed dependencies in one RUN layer
RUN apt-get update && \
    apt-get install -y curl ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/bin/indexserver /indexserver
COPY --from=builder /app/test/data /index

ENTRYPOINT ["/indexserver"]
