FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o push-swap ./cmd/push-swap && \
    go build -o checker ./cmd/checker

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/push-swap .
COPY --from=builder /app/checker .

ENTRYPOINT ["./push-swap"]
