# Stage 1: Builder
FROM golang:1.23-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the application
# - CGO_ENABLED=0: Disable CGO for a static binary (no external dynamic library dependencies)
# - GOOS=linux: Target Linux OS
# - -ldflags="-w -s": Strip debug information (DWARF, symbol table) to reduce binary size
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/main.go

# Stage 2: Runner
FROM alpine:latest

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/main .
RUN chown appuser:appgroup /app/main

USER appuser

EXPOSE 3000

CMD ["./main"]
