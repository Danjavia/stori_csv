# Base image for Go applications
FROM golang:1.19.2-bullseye AS builder

# Set working directory to project root
WORKDIR /app

# Copy Go modules and source code
COPY go.mod go.sum ./
RUN go clean -modcache
RUN go mod download
RUN go mod tidy
COPY . .

# Build the application
RUN go build -o /go/bin/app ./cmd/main.go

# Use a minimal image for the final stage
FROM alpine:latest

# Copy the built executable
COPY --from=builder /go/bin/app /app

# Expose port (if the application listens on a port)
EXPOSE 8080

# Run the application
CMD ["/app"]



# # Use the official Go image as the base image
# FROM golang:1.19.2-bullseye

# # Set the working directory
# WORKDIR /app

# # Copy the Go Modules file for efficient caching
# COPY go.mod .

# # Download dependencies
# RUN go mod download

# # Copy the rest of the application code
# COPY . .

# # Build the Go application
# RUN go build -o ./cmd/main.go

# # Expose the port the application will listen on (if applicable)
# EXPOSE 8080

# # Define the command to run when the container starts
# CMD ["./my-go-app"]