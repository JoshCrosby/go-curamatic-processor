# Start from the official Go image to build our application.
FROM golang:1.22.1-alpine3.19 AS builder

# Set the Current Working Directory inside the container.
WORKDIR /app

# Copy go mod and sum files.
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed.
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container.
COPY . .

# Build the Go app.
RUN GOOS=linux go build -a -o /app/main ./cmd/gocuramaticprocessor

# Start a new stage from scratch.
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage.
COPY --from=builder /app/main .

# Command to run the executable.
CMD ["./main"]
