FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum for dependency management
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go server binary
RUN go build -o server ./cmd/server/main.go

# Set execute permission for the server binary
RUN chmod +x ./server

# Expose the required port
EXPOSE 8081

# Set the default command to run the server
CMD ["./server"]
