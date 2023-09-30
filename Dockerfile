# Use the official Go image as a base image
FROM golang:1.17 as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go app for a smaller and secure artifact
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

### Start a new stage from scratch ###
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/main .

# Expose port 3000 for the application
EXPOSE 3200

# Command to run the application
CMD ["./main"]
