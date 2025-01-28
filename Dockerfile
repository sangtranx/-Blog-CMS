# Sử dụng image base chính thức của Go 1.23.5
FROM golang:1.23.5-alpine AS builder


# Thiết lập thư mục làm việc
WORKDIR /app


# Copy mã nguồn vào container
COPY . .


# Tải các dependencies
RUN go mod download


# Build ứng dụng
RUN go build -o main .


# Sử dụng image nhỏ gọn để chạy ứng dụng
FROM alpine:latest


# Thiết lập thư mục làm việc
WORKDIR /app


# Copy file binary từ builder
COPY --from=builder /app/main .
COPY --from=builder /app/storages ./storages
COPY --from=builder /app/Config ./Config


# Expose port mà ứng dụng sẽ chạy
EXPOSE 8082


# Lệnh để chạy ứng dụng
CMD ["./main"]
