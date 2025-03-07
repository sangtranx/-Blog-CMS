FROM golang:1.23.5-alpine AS builder

WORKDIR /app

COPY . .

# Setting dependencies
RUN go mod download

RUN go build -o main .

# using a small image
FROM alpine:latest

WORKDIR /app

# Copy file binary from builder
COPY --from=builder /app/main .
COPY --from=builder /app/storages ./storages
COPY --from=builder /app/Config ./Config

# Expose port
EXPOSE 8082

CMD ["./main"]
