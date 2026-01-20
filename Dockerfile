# Stage 1: Build the Go application
FROM golang:1.24.3 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
# CGO_ENABLED=0 ensures a statically linked binary, which is important for minimal base images like 'scratch' or 'alpine'
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Create the final, minimal runtime image
FROM alpine:latest

# Install necessary certificates (if your app makes HTTPS calls)
RUN apk update --no-cache && apk add --no-cache ca-certificates

# Set the working directory in the final container
WORKDIR /root/

# Copy the built binary from the 'builder' stage to the final image
COPY --from=builder /app/main .

# Copy the schedule.json file
COPY --from=builder /app/schedule.json .

# Expose the port your application listens on
EXPOSE 8080

# Command to run the executable when the container starts
CMD ["./main"]
