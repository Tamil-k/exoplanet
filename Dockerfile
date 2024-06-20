# Use an official Golang runtime as a parent image
FROM golang:1.21 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the entire project directory into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o exoplanet-service .

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary executable from the builder stage
COPY --from=builder /app/exoplanet-service .

# Expose port 9090 to the outside world
EXPOSE 9000

# Command to run the executable
CMD ["./exoplanet-service"]
