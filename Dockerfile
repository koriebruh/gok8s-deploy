# Multistage Dockerfile for building a Go application

# Stage 1: Build the Go application
FROM golang:1.24.1 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

# Stage 2: Create a minimal image with the built application
FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /root/
COPY --from=builder /app/app .
ENV APP_ENV=production
EXPOSE 3000
CMD ["./app"]