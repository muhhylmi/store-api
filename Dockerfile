# Multi-stage build: Build binary
FROM golang:latest AS builder

# Set working directory
WORKDIR /app

# Copy file aplikasi ke dalam container
COPY . .

# Build binary aplikasi Golang
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app

# Stage kedua: Gunakan image Alpine Linux yang ringan
FROM alpine:latest

# Update dan install dependensi yang dibutuhkan
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /app

# Copy binary dari builder stage ke dalam container
COPY --from=builder /app/app .

# Expose port yang digunakan oleh aplikasi
EXPOSE 3000

# Command untuk menjalankan aplikasi
CMD ["./app"]