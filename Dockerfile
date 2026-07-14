FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux \
    go build -o server ./cmd/server


FROM alpine:latest

RUN apk add --no-cache \
    nodejs \
    npm

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8084

CMD ["./server"]