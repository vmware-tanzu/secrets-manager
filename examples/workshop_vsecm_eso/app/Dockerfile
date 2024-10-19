FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY main.go .

RUN go build -o webhook-server main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/webhook-server .

EXPOSE 8080

CMD ["./webhook-server"]
