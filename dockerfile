FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o lpnapi .

FROM alpine:latest

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/lpnapi .

RUN chown appuser:appgroup lpnapi

USER appuser

EXPOSE 8080

CMD ["./lpnapi"]
