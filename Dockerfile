# Build stage
FROM golang:1.25-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application with verbose output
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server main.go

# Verify the binary was created
RUN ls -lah /build/ && \
    test -f /build/server || (echo "Binary not found!" && exit 1)

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /build/server .

# Make it executable
RUN chmod +x ./server

EXPOSE 8080

CMD ["./server"]