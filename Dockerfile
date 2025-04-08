FROM golang:1.23-alpine3.20 AS builder

WORKDIR /build

COPY . .

# Сборка бинарного файла
RUN go build -o main ./cmd

FROM alpine:3.20

COPY --from=builder /build/main /app/main

CMD ["/app/main"]