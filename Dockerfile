# Stage 1: Build the Go application
FROM golang:1.22.5-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go mod and sum files, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go app
RUN go build -o idr_sop_app .

# Use a smaller image for running the app
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the compiled Go app from the builder stage
COPY --from=builder /app/idr_sop_app .

# Copy the .env file (useful if you have it in the container)
COPY .env .env

# Expose the port (match the one specified in the .env file)
EXPOSE 8080

# Run the Go app
CMD ["./idr_sop_app"]

