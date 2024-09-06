# Build stage
FROM golang:1.22-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./main.go

# Start a new stage from scratch for the runtime environment
FROM alpine:latest AS runtime

# Install ca-certificates in case you need to make calls to HTTPS endpoints
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file and any other necessary files from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/config.yaml .
COPY --from=builder /app/.env .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]